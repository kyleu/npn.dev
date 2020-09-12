package call

import (
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
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
	bod, err := body.Parse(headers.GetValue("Content-Type"), r.Body)
	var es *string = nil
	if err != nil {
		ex := err.Error()
		es = &ex
	}
	return &Response{
		Status:           r.Status,
		StatusCode:       r.StatusCode,
		Proto:            r.Proto,
		ProtoMajor:       r.ProtoMajor,
		ProtoMinor:       r.ProtoMinor,
		Headers:          headers,
		ContentLength:    r.ContentLength,
		TransferEncoding: r.TransferEncoding,
		Close:            r.Close,
		Uncompressed:     r.Uncompressed,
		Body:             bod,
		Error:            es,
	}
}
