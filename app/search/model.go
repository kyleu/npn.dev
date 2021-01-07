package search

type Params struct {
	Q      string `json:"q"`
	Offset int    `json:"o,omitempty"`
}

type Result struct {
	T        string `json:"t"`
	Key      string `json:"key"`
	Prelude  string `json:"prelude"`
	Postlude string `json:"postlude"`
	Loc      string `json:"loc"`
}

type Results []*Result
