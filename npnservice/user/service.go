package user

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
)

// Interface for managing SystemUser instances
type Service interface {
	// Lists SystemUser instances matching the provided params
	List(params *npncore.Params) SystemUsers
	// Returns the SystemUser matching the provided id
	GetByID(userID uuid.UUID, addIfMissing bool) *SystemUser
	// Saves an npnuser.UserProfile to storage
	SaveProfile(prof *npnuser.UserProfile) (*npnuser.UserProfile, error)
	// Updates the name and picture of the SystemUser matching the provided id
	UpdateMember(userID uuid.UUID, name string, picture string) error
	// Sets the role of the SystemUser matching the provided id
	SetRole(userID uuid.UUID, role npnuser.Role) error
}
