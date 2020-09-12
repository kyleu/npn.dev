package call

import (
	"github.com/kyleu/npn/app/request/header"
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
	// TLS *tls.ConnectionState `json:"tls,omitempty"`
	// Body []byte maybe? `json:"body,omitempty"`
}

func ResponseFromHTTP(r *http.Response) *Response {
	headers := make(header.Headers, 0, len(r.Header))
	for k, vs := range r.Header {
		for _, v := range vs {
			headers = append(headers, &header.Header{Key: k, Value: v})
		}
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
	}
}
