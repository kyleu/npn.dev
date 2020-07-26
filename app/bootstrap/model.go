package bootstrap

import "encoding/json"

type Prototype struct {
	Key         string   `json:"key"`
	Description string   `json:"description,omitempty"`
	Inherit     []string `json:"inherit,omitempty"`
	Folders     []string `json:"folders,omitempty"`
	BuildCmds   []string `json:"buildCmds,omitempty"`
}

func (p *Prototype) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Key)
}

func (p *Prototype) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	p = PrototypeFromString(s)
	return nil
}

var (
	Unmanaged = &Prototype{
		Key:         "unmanaged",
		Description: "simple project with no build process",
	}
)

var GoSimple = &Prototype{
	Key:         "go-simple",
	Description: "bare-bones golang project",
	Folders:     []string{"go/simple"},
	BuildCmds:   []string{"goimports -w .", "make build", "build/{{.Key}}"},
}

var GoDatabase = &Prototype{
	Key:         "go-database",
	Description: "simple golang project with database support",
	Folders:     []string{"go/database"},
	BuildCmds:   []string{"goimports -w .", "make build", "build/{{.Key}}"},
}

var AllPrototypes = []*Prototype{Unmanaged, GoSimple, GoDatabase}

func PrototypeFromString(s string) *Prototype {
	for _, t := range AllPrototypes {
		if t.Key == s {
			return t
		}
	}
	return Unmanaged
}
