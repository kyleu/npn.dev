package request

type Summary struct {
	Key         string `json:"key,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Order       int    `json:"order,omitempty"`
}

func (r *Summary) TitleWithFallback() string {
	if len(r.Title) == 0 {
		return r.Key
	}
	return r.Title
}

type Summaries []*Summary
