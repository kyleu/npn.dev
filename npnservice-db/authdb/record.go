package authdb

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/npn/npnuser"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/gofrs/uuid"
)

func (s *ServiceDatabase) NewRecord(r *auth.Record) (*auth.Record, error) {
	q := npndatabase.SQLInsert(npncore.KeyAuth, []string{npncore.KeyID, npncore.WithDBID(npncore.KeyUser), npncore.KeyProvider, npncore.WithDBID(npncore.KeyProvider), "user_list_id", "user_list_name", "access_token", "expires", npncore.KeyName, npncore.KeyEmail, "picture"}, 1)
	err := s.db.Insert(q, nil, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.UserID, r.ID), nil
}

func (s *ServiceDatabase) UpdateRecord(r *auth.Record) error {
	cols := []string{"user_list_id", "user_list_name", "access_token", "expires", npncore.KeyName, npncore.KeyEmail, "picture"}
	q := npndatabase.SQLUpdate(npncore.KeyAuth, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	return s.db.UpdateOne(q, nil, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture, r.ID)
}

func (s *ServiceDatabase) MergeProfile(p *npnuser.UserProfile, record *auth.Record) (*auth.Record, error) {
	p.Name = record.Name
	if len(p.Name) == 0 {
		p.Name = record.Provider.Title + " User"
	}
	p.Picture = record.Picture

	_, err := s.users.SaveProfile(p)
	if err != nil {
		return nil, errors.Wrap(err, "error saving user profile")
	}

	return record, nil
}

func (s *ServiceDatabase) GetByID(userID uuid.UUID, authID uuid.UUID) *auth.Record {
	dto := &recordDTO{}
	q := npndatabase.SQLSelectSimple("*", npncore.KeyAuth, npncore.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, authID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting auth record by id [%v]: %+v", authID, err))
		return nil
	}
	return dto.toRecord()
}

func (s *ServiceDatabase) GetByProviderID(userID uuid.UUID, key string, code string) *auth.Record {
	dto := &recordDTO{}
	q := npndatabase.SQLSelectSimple("*", npncore.KeyAuth, "provider = $1 and provider_id = $2")
	err := s.db.Get(dto, q, nil, key, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting auth record by provider [%v:%v]: %+v", key, code, err))
		return nil
	}
	return dto.toRecord()
}

func (s *ServiceDatabase) GetByUserID(userID uuid.UUID, params *npncore.Params) auth.Records {
	if s.db == nil {
		return nil
	}

	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAuth, params, npncore.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := npndatabase.SQLSelect("*", npncore.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving auths for user [%v]: %+v", userID, err))
		return nil
	}
	return toRecords(dtos)
}

func (s *ServiceDatabase) Delete(userID uuid.UUID, authID uuid.UUID) error {
	q := npndatabase.SQLDelete(npncore.KeyAuth, npncore.KeyID+" = $1")
	return s.db.DeleteOne(q, nil, authID)
}

func toRecords(dtos []recordDTO) auth.Records {
	ret := make(auth.Records, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toRecord())
	}
	return ret
}
