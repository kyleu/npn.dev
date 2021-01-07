package collection

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"strings"
)

type Collection struct {
	Key          string   `json:"key"`
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Path         string   `json:"path,omitempty"`
	RequestOrder []string `json:"requestOrder,omitempty"`
}

var defaultCollection = &Collection{Key: "_", Title: "Default"}

func (c *Collection) TitleWithFallback() string {
	if len(c.Title) == 0 {
		return c.Key
	}
	return c.Title
}

func (c *Collection) Normalize(key string, p string) *Collection {
	c.Key = key
	c.Path = p
	return c
}

func check(s string, q string) (bool, string, string) {
	low := strings.ToLower(s)
	if strings.Contains(low, q) {
		idx := strings.Index(low, q)
		return true, s[0:idx], s[idx + len(q):]
	}
	return false, "", ""
}

func (c *Collection) Matches(q string) (bool, string, string, string) {
	q = strings.ToLower(q)

	matched, pre, post := check(c.Key, q)
	if matched {
		return true, pre, "key", post
	}
	matched, pre, post = check(c.Title, q)
	if matched {
		return true, pre, "title", post
	}
	matched, pre, post = check(c.Description, q)
	if matched {
		return true, pre, "description", post
	}
	return false, "", "", ""
}

type Collections []*Collection

type Summary struct {
	Coll  *Collection `json:"coll"`
	Count int         `json:"count"`
}

type Summaries []*Summary

type FullCollection struct {
	Coll     *Collection      `json:"coll"`
	Requests request.Requests `json:"reqs"`
	Sess     *session.Session `json:"sess"`
}
