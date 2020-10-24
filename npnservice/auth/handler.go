package auth

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npnuser"
)

func Handle(s Service, profile *npnuser.UserProfile, prv *Provider, code string) (*Record, error) {
	if !s.Enabled() {
		return nil, ErrorAuthDisabled
	}

	if profile == nil {
		return nil, errors.New("no user profile for auth")
	}

	cfg := GetConfig(s, prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	record, err := DecodeRecord(s, prv, code)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving auth profile")
	}
	if record == nil {
		return nil, errors.New("cannot retrieve auth profile")
	}
	record.UserID = profile.UserID

	curr := s.GetByProviderID(record.UserID, record.Provider.Key, record.ProviderID)
	if curr == nil {
		record, err = s.NewRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new auth record")
		}

		return s.MergeProfile(profile, record)
	}
	if curr.UserID == profile.UserID {
		record.ID = curr.ID

		err = s.UpdateRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error updating auth record")
		}

		return s.MergeProfile(profile, record)
	}

	record, err = s.NewRecord(record)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new auth record")
	}

	return s.MergeProfile(profile, record)
}
