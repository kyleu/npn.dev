package userdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/npn/npnservice/user"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

type ServiceDatabase struct {
	db     *npndatabase.Service
	logger logur.Logger
}

var _ user.Service = (*ServiceDatabase)(nil)

func NewServiceDatabase(db *npndatabase.Service, logger logur.Logger) user.Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyUser})
	return &ServiceDatabase{db: db, logger: logger}
}

const userTable = "system_user"

func (s *ServiceDatabase) new(id uuid.UUID) (*user.SystemUser, error) {
	s.logger.Info("creating user [" + id.String() + "]")

	q := npndatabase.SQLInsert(userTable, []string{npncore.KeyID, npncore.KeyName, npncore.KeyRole, npncore.KeyTheme, "nav_color", "link_color", "picture", "locale", npncore.KeyCreated}, 1)
	prof := npnuser.NewUserProfile(id, "")
	err := s.db.Insert(q, nil, prof.UserID, prof.Name, npnuser.RoleGuest.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), time.Now())

	if err != nil {
		return nil, err
	}

	return s.GetByID(id, false), nil
}

func (s *ServiceDatabase) List(params *npncore.Params) user.SystemUsers {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyUser, params, npncore.DefaultCreatedOrdering...)

	var ret user.SystemUsers

	q := npndatabase.SQLSelect("*", userTable, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&ret, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving system users: %+v", err))
		return nil
	}

	return ret
}

func (s *ServiceDatabase) GetByID(id uuid.UUID, addIfMissing bool) *user.SystemUser {
	ret := &user.SystemUser{}
	q := npndatabase.SQLSelectSimple("*", userTable, npncore.KeyID+" = $1")
	err := s.db.Get(ret, q, nil, id)
	if err == sql.ErrNoRows {
		if addIfMissing {
			ret, err := s.new(id)
			if err != nil {
				s.logger.Error(fmt.Sprintf("error creating new user with id [%v]: %+v", id, err))
			}
			return ret
		}
		return nil
	}
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting user by id [%v]: %+v", id, err))
		return nil
	}
	return ret
}

func (s *ServiceDatabase) GetByCreated(d *time.Time, params *npncore.Params) user.SystemUsers {
	params = npncore.ParamsWithDefaultOrdering(userTable, params, npncore.DefaultCreatedOrdering...)
	var ret user.SystemUsers
	q := npndatabase.SQLSelect("*", userTable, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&ret, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving users created on [%v]: %+v", d, err))
		return nil
	}
	return ret
}

func (s *ServiceDatabase) SaveProfile(prof *npnuser.UserProfile) (*npnuser.UserProfile, error) {
	s.logger.Debug("updating user [" + prof.UserID.String() + "] from profile")
	cols := []string{"name", "role", "theme", "nav_color", "link_color", "picture", "locale"}
	q := npndatabase.SQLUpdate(userTable, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, prof.Name, prof.Role.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), prof.UserID)
	if err != nil {
		return nil, err
	}
	return prof, nil
}

func (s *ServiceDatabase) UpdateMember(userID uuid.UUID, name string, picture string) error {
	s.logger.Debug("updating user [" + userID.String() + "]")
	cols := []string{"name", "picture"}
	q := npndatabase.SQLUpdate(userTable, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, name, picture, userID)
	return err
}

func (s *ServiceDatabase) SetRole(userID uuid.UUID, role npnuser.Role) error {
	_ = s.GetByID(userID, true)
	s.logger.Info("updating user role [" + userID.String() + "]")
	cols := []string{"role"}
	q := npndatabase.SQLUpdate(userTable, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, role.Key, userID)
	return err
}
