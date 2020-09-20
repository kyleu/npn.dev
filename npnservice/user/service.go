package user

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
)

type Service interface {
	List(params *npncore.Params) SystemUsers
	GetByID(userID uuid.UUID, addIfMissing bool) *SystemUser
	SaveProfile(prof *npnuser.UserProfile) (*npnuser.UserProfile, error)
	UpdateMember(userID uuid.UUID, name string, picture string) error
	SetRole(userID uuid.UUID, role npnuser.Role) error
}
