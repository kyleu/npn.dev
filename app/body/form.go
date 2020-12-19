package body

import (
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
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

func (f *FormEntry) Clone() *FormEntry {
	return &FormEntry{K: f.K, V: f.V}
}

type FormData []*FormEntry

func (f FormData) String() string {
	ret := make([]string, 0, len(f))
	for _, x := range f {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, "&")
}

func (f FormData) Clone() FormData {
	ret := make(FormData, len(f))
	for _, x := range ret {
		ret = append(ret, x.Clone())
	}
	return ret
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

func (f *Form) Merge(data npncore.Data, logger logur.Logger) Config {
	d := make(FormData, 0, len(f.Data))
	for _, dt := range f.Data {
		d = append(d, &FormEntry{
			K: npncore.MergeLog("body.form.data.k", dt.K, data, logger),
			V: npncore.MergeLog("body.form.data.v", dt.V, data, logger),
		})
	}

	s := f.str
	if len(s) > 0 {
		s = npncore.MergeLog("body.form.str", s, data, logger)
	}

	return &Form{
		Data: d,
		str:  s,
	}
}

func (f *Form) Clone() *Body {
	return NewBody(KeyForm, &Form{Data: f.Data.Clone(), str: f.str})
}
