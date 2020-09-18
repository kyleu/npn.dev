package call

import (
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/npncore"
	"net/http"
)

type Response struct {
	Status           string         `json:"status,omitempty"`
	StatusCode       int            `json:"statusCode,omitempty"`
	Proto            string         `json:"proto,omitempty"`
	ProtoMajor       int            `json:"protoMajor,omitempty"`
	ProtoMinor       int            `json:"protoMinor,omitempty"`
	Headers          header.Headers `json:"headers,omitempty"`
	ContentLength    int64          `json:"contentLength,omitempty"`
	ContentType      string         `json:"contentType,omitempty"`
	Charset          string         `json:"charset,omitempty"`
	TransferEncoding []string       `json:"transferEncoding,omitempty"`
	Close            bool           `json:"close,omitempty"`
	Uncompressed     bool           `json:"uncompressed,omitempty"`
	Body             *body.Body     `json:"body,omitempty"`
	Error            *string        `json:"error,omitempty"`
}

func ResponseFromHTTP(r *http.Response) *Response {
	headers := make(header.Headers, 0, len(r.Header))
	for k, vs := range r.Header {
		for _, v := range vs {
			headers = append(headers, &header.Header{Key: k, Value: v})
		}
	}
	ct, charset := parseCT(headers.GetValue("Content-Type"))
	bod, err := body.Parse(ct, r.ContentLength, r.Body)
	var es *string = nil
	if err != nil {
		ex := err.Error()
		es = &ex
	}
	if len(ct) == 0 && bod != nil && bod.Config != nil {
		ct = bod.Config.MimeType()
	}
	return &Response{
		Status:           r.Status,
		StatusCode:       r.StatusCode,
		Proto:            r.Proto,
		ProtoMajor:       r.ProtoMajor,
		ProtoMinor:       r.ProtoMinor,
		Headers:          headers,
		ContentLength:    r.ContentLength,
		ContentType:      ct,
		Charset:          charset,
		TransferEncoding: r.TransferEncoding,
		Close:            r.Close,
		Uncompressed:     r.Uncompressed,
		Body:             bod,
		Error:            es,
	}
}

func parseCT(h string) (string, string) {
	return npncore.SplitString(h, ';', true)
}
