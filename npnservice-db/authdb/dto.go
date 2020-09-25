package authdb

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/auth"
	"time"
)

type recordDTO struct {
	ID           uuid.UUID  `db:"id"`
	UserID       uuid.UUID  `db:"user_id"`
	Provider     string     `db:"provider"`
	ProviderID   string     `db:"provider_id"`
	UserListID   string     `db:"user_list_id"`
	UserListName string     `db:"user_list_name"`
	AccessToken  string     `db:"access_token"`
	Expires      *time.Time `db:"expires"`
	Name         string     `db:"name"`
	Email        string     `db:"email"`
	Picture      string     `db:"picture"`
	Created      time.Time  `db:"created"`
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

