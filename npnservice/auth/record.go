package auth

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"time"

	"github.com/gofrs/uuid"
)

func (s *Service) NewRecord(r *Record) (*Record, error) {
	q := npndatabase.SQLInsert(npncore.KeyAuth, []string{npncore.KeyID, npncore.WithDBID(npncore.KeyUser), npncore.KeyProvider, npncore.WithDBID(npncore.KeyProvider), "user_list_id", "user_list_name", "access_token", "expires", npncore.KeyName, npncore.KeyEmail, "picture"}, 1)
	err := s.db.Insert(q, nil, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID), nil
}

func (s *Service) UpdateRecord(r *Record) error {
	cols := []string{"user_list_id", "user_list_name", "access_token", "expires", npncore.KeyName, npncore.KeyEmail, "picture"}
	q := npndatabase.SQLUpdate(npncore.KeyAuth, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	return s.db.UpdateOne(q, nil, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture, r.ID)
}

func (s *Service) List(params *npncore.Params) Records {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAuth, params, npncore.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := npndatabase.SQLSelect("*", npncore.KeyAuth, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving auth records: %+v", err))
		return nil
	}
	return toRecords(dtos)
}

func (s *Service) GetByCreated(d *time.Time, params *npncore.Params) Records {
	params = npncore.ParamsWithDefaultOrdering("system_user", params, npncore.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := npndatabase.SQLSelect("*", "system_user", "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving auth records created on [%v]: %+v", d, err))
		return nil
	}
	return toRecords(dtos)
}

func (s *Service) GetByID(authID uuid.UUID) *Record {
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

func (s *Service) GetByProviderID(key string, code string) *Record {
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

func (s *Service) GetByUserID(userID uuid.UUID, params *npncore.Params) Records {
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

func (s *Service) Delete(authID uuid.UUID) error {
	q := npndatabase.SQLDelete(npncore.KeyAuth, npncore.KeyID+" = $1")
	return s.db.DeleteOne(q, nil, authID)
}

func toRecords(dtos []recordDTO) Records {
	ret := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toRecord())
	}
	return ret
}
