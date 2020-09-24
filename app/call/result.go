package call

type Result struct {
	Status   string    `json:"status,omitempty"`
	Response *Response `json:"response,omitempty"`
	Timing   *Timing   `json:"timing,omitempty"`
	Error    string    `json:"error,omitempty"`
}
