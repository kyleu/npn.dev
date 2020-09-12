package request

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Method struct {
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
}

var (
	MethodGet     = Method{Key: http.MethodGet, Description: ""}
	MethodHead    = Method{Key: http.MethodHead, Description: ""}
	MethodPost    = Method{Key: http.MethodPost, Description: ""}
	MethodPut     = Method{Key: http.MethodPut, Description: ""}
	MethodPatch   = Method{Key: http.MethodPatch, Description: ""}
	MethodDelete  = Method{Key: http.MethodDelete, Description: ""}
	MethodConnect = Method{Key: http.MethodConnect, Description: ""}
	MethodOptions = Method{Key: http.MethodOptions, Description: ""}
	MethodTrace   = Method{Key: http.MethodTrace, Description: ""}
)
var AllMethods = []Method{MethodGet, MethodHead, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodConnect, MethodOptions, MethodTrace}

func MethodFromString(s string) Method {
	s = strings.ToUpper(s)
	for _, t := range AllMethods {
		if t.Key == s {
			return t
		}
	}
	return Method{Key: s, Description: "Custom method"}
}

func (t *Method) String() string {
	return t.Key
}

func (t *Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Method) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = MethodFromString(s)
	return nil
}
