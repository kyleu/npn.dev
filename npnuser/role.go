package npnuser

import (
	"encoding/json"
)

type Role struct {
	Key string
}

var RoleGuest = Role{
	Key: "guest",
}

var RoleUser = Role{
	Key: "user",
}

var RoleAdmin = Role{
	Key: "admin",
}

var AllRoles = []Role{RoleGuest, RoleUser, RoleAdmin}

func RoleFromString(s string) Role {
	for _, t := range AllRoles {
		if t.Key == s {
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
