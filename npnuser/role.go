package npnuser

import (
	"encoding/json"
)

// A user's Role
type Role struct {
	Key string
}

// Restricted default role
var RoleGuest = Role{
	Key: "guest",
}

// Regular user of the system
var RoleUser = Role{
	Key: "user",
}

// An adminstrator of the system
var RoleAdmin = Role{
	Key: "admin",
}

var AllRoles = []Role{RoleGuest, RoleUser, RoleAdmin}

// Finds the Role matching the provided key, or RoleGuest
func RoleFromString(key string) Role {
	for _, t := range AllRoles {
		if t.Key == key {
			return t
		}
	}
	return RoleGuest
}

func (t *Role) String() string {
	return t.Key
}

func (t *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Role) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = RoleFromString(s)
	return nil
}
