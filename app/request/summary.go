package request

import "strings"

type Summary struct {
	Key         string `json:"key,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Order       int    `json:"order,omitempty"`
}

func (r *Summary) TitleWithFallback() string {
	if r.Title == "" {
		return r.Key
	}
	return r.Title
}

func check(s string, q string) (bool, string, string) {
	low := strings.ToLower(s)
	if strings.Contains(low, q) {
		idx := strings.Index(low, q)
		return true, s[0:idx], s[idx+len(q):]
	}
	return false, "", ""
}

func (r *Summary) Matches(q string) (bool, string, string, string) {
	q = strings.ToLower(q)

	matched, pre, post := check(r.Key, q)
	if matched {
		return true, pre, "key", post
	}
	matched, pre, post = check(r.Title, q)
	if matched {
		return true, pre, "title", post
	}
	matched, pre, post = check(r.Description, q)
	if matched {
		return true, pre, "description", post
	}
	matched, pre, post = check(r.URL, q)
	if matched {
		return true, pre, "url", post
	}
	return false, "", "", ""
}

type Summaries []*Summary
