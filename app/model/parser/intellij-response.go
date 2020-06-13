package parser

import "github.com/kyleu/npn/app/model/schema"

type IntelliJResult interface {
	ParentID() int
}

type IntelliJResponse struct {
	DSN      string
	DBType   string
	DBFamily string
	Data     []IntelliJResult
	ByParent map[int][]IntelliJResult
}

func NewIntelliJResponse() *IntelliJResponse {
	return &IntelliJResponse{
		Data:     make([]IntelliJResult, 0),
		ByParent: make(map[int][]IntelliJResult),
	}
}

func (r *IntelliJResponse) Schema() (*schema.Schema, interface{}, error) {
	ret := &schema.Schema{}
	return ret, r, nil
}
