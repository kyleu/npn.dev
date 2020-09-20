package imprt

type Output struct {
	Filename string      `json:"filename,omitempty"`
	Type     string      `json:"type,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Error    string      `json:"error,omitempty"`
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

type phase struct {
	Key   string
	Value interface{}
	Error error
	Final bool
}

func errorPhase(err error, x interface{}) *phase {
	return &phase{Key: "error", Value: x, Error: err, Final: true}
}
