package collection

type Collection struct {
	Key         string `json:"-"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Path        string `json:"-"`
}

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
