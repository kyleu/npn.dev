package util

import (
	"encoding/json"
)

type Theme struct {
	Name string
	CSS  string
}

var ThemeAuto = Theme{
	Name: "auto",
	CSS:  "uk-dark",
}

var ThemeLight = Theme{
	Name: "light",
	CSS:  "uk-dark",
}

var ThemeDark = Theme{
	Name: "dark",
	CSS:  "uk-light",
}

var AllThemes = []Theme{ThemeAuto, ThemeLight, ThemeDark}

func ThemeFromString(s string) Theme {
	for _, t := range AllThemes {
		if t.Name == s {
			return t
		}
	}
	return ThemeAuto
}

func (t *Theme) String() string {
	return t.Name
}

func (t *Theme) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}

func (t *Theme) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = ThemeFromString(s)
	return nil
}

var AllColors = []string{"clear", "grey", "bluegrey", "red", "orange", "yellow", "green", "blue", "purple"}
