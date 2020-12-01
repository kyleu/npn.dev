package userfs

import (
	"fmt"
	"time"

	"github.com/kyleu/npn/npnservice/user"

	"golang.org/x/text/language"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

var systemUser = &user.SystemUser{
	UserID:   npnuser.SystemUserID,
	Name:     "Guest",
	Role:     "admin",
	Settings: npnuser.DefaultSettings,
	Picture:  "",
	Locale:   language.AmericanEnglish.String(),
	Created:  time.Time{},
}

type ServiceFilesystem struct {
	Multiuser bool
	files     npncore.FileLoader
	logger    logur.Logger
}

var _ user.Service = (*ServiceFilesystem)(nil)

func NewServiceFilesystem(multiuser bool, files npncore.FileLoader, logger logur.Logger) *ServiceFilesystem {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyUser})
	return &ServiceFilesystem{Multiuser: multiuser, files: files, logger: logger}
}

func (s *ServiceFilesystem) new(userID uuid.UUID) (*user.SystemUser, error) {
	s.logger.Info("creating user [" + userID.String() + "]")
	np := npnuser.NewUserProfile(userID, "Guest")
	p, err := s.SaveProfile(np)
	return user.FromProfile(p, time.Now()), err
}

func (s *ServiceFilesystem) List(params *npncore.Params) user.SystemUsers {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyUser, params, npncore.DefaultCreatedOrdering...)

	var ret user.SystemUsers

	if s.Multiuser {
		x := s.files.ListDirectories("/users")
		for _, f := range x {
			uid := npncore.GetUUIDFromString(f)
			ret = append(ret, s.GetByID(*uid, false))
		}
	} else {
		ret = append(ret, systemUser)
	}

	return ret
}

func (s *ServiceFilesystem) GetByID(userID uuid.UUID, addIfMissing bool) *user.SystemUser {
	tgt := &npnuser.UserProfile{}
	fn := s.filenameFor(userID)
	exists, _ := s.files.Exists(fn)
	if !exists && addIfMissing {
		_, err := s.new(userID)
		if err != nil {
			s.logger.Warn(fmt.Sprintf("can't save new profile: %+v", err))
			return nil
		}
	}

	content, err := s.files.ReadFile(fn)
	if err == nil {
		err = npncore.FromJSON(content, tgt)
		if err != nil {
			s.logger.Warn(fmt.Sprintf("can't load profile: %+v", err))
			return nil
		}
	} else {
		return nil
	}
	return user.FromProfile(tgt, time.Now())
}

func (s *ServiceFilesystem) SaveProfile(prof *npnuser.UserProfile) (*npnuser.UserProfile, error) {
	err := s.files.WriteFile(s.filenameFor(prof.UserID), npncore.ToJSONBytes(prof.ToProfile(), s.logger, true), true)
	if err != nil {
		return nil, err
	}
	return prof, nil
}

func (s *ServiceFilesystem) UpdateMember(userID uuid.UUID, name string, picture string) error {
	su := s.GetByID(userID, true)
	su.Name = name
	su.Picture = picture
	_, err := s.SaveProfile(su.ToProfile())
	return err
}

func (s *ServiceFilesystem) SetRole(userID uuid.UUID, role npnuser.Role) error {
	su := s.GetByID(userID, true)
	su.Role = role.String()
	_, err := s.SaveProfile(su.ToProfile())
	return err
}

func (s *ServiceFilesystem) filenameFor(id uuid.UUID) string {
	if (!s.Multiuser) || id == systemUser.UserID {
		return "profile.json"
	}
	return "/users/" + id.String() + "/profile.json"
}
