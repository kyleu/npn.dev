package call

type Result struct {
	Status   string    `json:"status,omitempty"`
	Response *Response `json:"response,omitempty"`
	Duration int       `json:"duration,omitempty"`
	Error    *string   `json:"error,omitempty"`
}
