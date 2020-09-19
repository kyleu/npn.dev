package imprt

type Output struct {
	Filename string      `json:"filename,omitempty"`
	Type     string      `json:"type,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

type Outputs []*Output

type File struct {
	Filename    string `json:"filename,omitempty"`
	Size        int64  `json:"size,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}

type Config struct {
	Files  []File `json:"files,omitempty"`
	Status string `json:"status,omitempty"`
}
