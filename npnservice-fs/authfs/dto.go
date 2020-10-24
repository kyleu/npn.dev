package authfs

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/auth"
)

type recordDTO struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	Provider     string     `json:"provider"`
	ProviderID   string     `json:"provider_id"`
	UserListID   string     `json:"user_list_id"`
	UserListName string     `json:"user_list_name"`
	AccessToken  string     `json:"access_token"`
	Expires      *time.Time `json:"expires"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Picture      string     `json:"picture"`
	Created      time.Time  `json:"created"`
}

func newDTO(a *auth.Record) *recordDTO {
	return &recordDTO{
		ID:           a.ID,
		UserID:       a.UserID,
		Provider:     a.Provider.Key,
		ProviderID:   a.ProviderID,
		UserListID:   a.UserListID,
		UserListName: a.UserListName,
		AccessToken:  a.AccessToken,
		Expires:      a.Expires,
		Name:         a.Name,
		Email:        a.Email,
		Picture:      a.Picture,
		Created:      a.Created,
	}
}

func (dto *recordDTO) toRecord() *auth.Record {
	return &auth.Record{
		ID:           dto.ID,
		UserID:       dto.UserID,
		Provider:     auth.ProviderFromString(dto.Provider),
		ProviderID:   dto.ProviderID,
		UserListID:   dto.UserListID,
		UserListName: dto.UserListName,
		AccessToken:  dto.AccessToken,
		Expires:      dto.Expires,
		Name:         dto.Name,
		Email:        dto.Email,
		Picture:      dto.Picture,
		Created:      dto.Created,
	}
}
