package authfs

import (
	"fmt"
	"path"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnuser"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"
)

func (s *ServiceFS) NewRecord(r *auth.Record) (*auth.Record, error) {
	return s.write(r)
}

func (s *ServiceFS) UpdateRecord(r *auth.Record) error {
	_, err := s.write(r)
	return err
}

func (s *ServiceFS) MergeProfile(p *npnuser.UserProfile, record *auth.Record) (*auth.Record, error) {
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

func (s *ServiceFS) GetByID(userID uuid.UUID, authID uuid.UUID) *auth.Record {
	ud := userDir(userID)
	fn := path.Join(ud, authID.String()+".json")
	exists, _ := s.files.Exists(fn)
	if !exists {
		return nil
	}

	content, err := s.files.ReadFile(fn)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("can't load auth [%v]: %+v", authID, err))
		return nil
	}

	dto := &recordDTO{}
	err = npncore.FromJSON(content, dto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting auth record by id [%v]: %+v", authID, err))
		return nil
	}
	return dto.toRecord()
}

func (s *ServiceFS) GetByProviderID(userID uuid.UUID, key string, code string) *auth.Record {
	curr := s.GetByUserID(userID, nil)
	for _, r := range curr {
		if (r.Provider != nil && r.Provider.Key == key) && (r.ProviderID == code) {
			return r
		}
	}
	return nil
}

func (s *ServiceFS) GetByUserID(userID uuid.UUID, params *npncore.Params) auth.Records {
	if !s.enabled {
		return nil
	}

	ud := userDir(userID)
	curr := s.files.ListJSON(ud)
	if len(curr) == 0 {
		return nil
	}

	ret := make(auth.Records, 0, len(curr))
	for _, authID := range curr {
		fn := path.Join(ud, authID+".json")
		content, err := s.files.ReadFile(fn)

		dto := &recordDTO{}
		err = npncore.FromJSON(content, dto)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error getting auth record by id [%v]: %+v", authID, err))
			return nil
		}
		ret = append(ret, dto.toRecord())
	}

	return ret
}

func (s *ServiceFS) Delete(userID uuid.UUID, authID uuid.UUID) error {
	ud := userDir(userID)
	fn := path.Join(ud, authID.String()+".json")
	return s.files.Remove(fn)
}

func toRecords(dtos []recordDTO) auth.Records {
	ret := make(auth.Records, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toRecord())
	}
	return ret
}

func (s *ServiceFS) write(r *auth.Record) (*auth.Record, error) {
	ud := userDir(r.UserID)
	err := s.files.CreateDirectory(ud)
	if err != nil {
		return nil, err
	}

	ar := newDTO(r)

	js := npncore.ToJSONBytes(ar, s.logger, true)
	err = s.files.WriteFile(path.Join(ud, r.ID.String()+".json"), js, true)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.UserID, r.ID), nil
}

func userDir(u uuid.UUID) string {
	return path.Join("users", u.String(), "auth")
}
