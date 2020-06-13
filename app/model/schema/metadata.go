package schema

type Metadata struct {
	Comments []string `json:"comments,omitempty"`
	Origin Origin `json:"origin,omitempty"`
	Source string `json:"source,omitempty"`
	Line int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}
