package authdb

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnuser"
)

func (s *ServiceDatabase) Handle(profile *npnuser.UserProfile, prv *auth.Provider, code string) (*auth.Record, error) {
	if !s.Enabled() {
		return nil, auth.ErrorAuthDisabled
	}

	if profile == nil {
		return nil, errors.New("no user profile for auth")
	}

	cfg := s.getConfig(prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	record, err := auth.DecodeRecord(s, prv, code)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving auth profile")
	}
	if record == nil {
		return nil, errors.New("cannot retrieve auth profile")
	}
	record.UserID = profile.UserID

	curr := s.GetByProviderID(record.Provider.Key, record.ProviderID)
	if curr == nil {
		record, err = s.NewRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new auth record")
		}

		return s.mergeProfile(profile, record)
	}
	if curr.UserID == profile.UserID {
		record.ID = curr.ID

		err = s.UpdateRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error updating auth record")
		}

		return s.mergeProfile(profile, record)
	}

	record, err = s.NewRecord(record)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new auth record")
	}

	return s.mergeProfile(profile, record)
}

func (s *ServiceDatabase) mergeProfile(p *npnuser.UserProfile, record *auth.Record) (*auth.Record, error) {
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
