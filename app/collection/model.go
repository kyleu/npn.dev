package collection

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
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
