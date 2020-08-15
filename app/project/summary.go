package project

type Summary struct {
	Key         string   `json:"key"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	SchemaKeys  []string `json:"schemaKeys,omitempty"`
}

type Summaries []*Summary
