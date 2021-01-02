package search

type Params struct {
	Q string `json:"q"`
}

type Result struct {
	Msg string `json:"msg"`
}

type Results []*Result
