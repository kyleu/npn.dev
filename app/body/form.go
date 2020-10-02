package body

import (
	"net/url"
	"strings"
)

const KeyForm = "form"

type FormEntry struct {
	K string `json:"k"`
	V string `json:"v"`
}

func (f *FormEntry) String() string {
	return url.QueryEscape(f.K) + "=" + url.QueryEscape(f.V)
}

type FormData []*FormEntry

func (f FormData) String() string {
	ret := make([]string, 0, len(f))
	for _, x := range f {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, "&")
}

type Form struct {
	Data FormData `json:"data"`
	str  string
}

var _ Config = (*Form)(nil)

func NewForm(data ...*FormEntry) *Body {
	return NewBody(KeyForm, &Form{Data: data})
}

func (f *Form) ContentLength() int64 {
	return int64(len(f.Bytes()))
}

func (f *Form) Bytes() []byte {
	return []byte(f.String())
}

func (f *Form) MimeType() string {
	return "application/x-www-form-urlencoded"
}

func (f *Form) String() string {
	if len(f.str) == 0 {
		f.str = f.Data.String()
	}
	return f.str
}
