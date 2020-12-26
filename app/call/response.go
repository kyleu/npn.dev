package call

import (
	"net/http"
	"strings"

	"github.com/kyleu/npn/app/session"

	"github.com/kyleu/npn/app/request"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
)

type Response struct {
	Method           string         `json:"method,omitempty"`
	URL              string         `json:"url,omitempty"`
	RequestHeaders   header.Headers `json:"requestHeaders,omitempty"`
	Status           string         `json:"status"`
	StatusCode       int            `json:"statusCode,omitempty"`
	Proto            string         `json:"proto,omitempty"`
	ProtoMajor       int            `json:"protoMajor,omitempty"`
	ProtoMinor       int            `json:"protoMinor,omitempty"`
	Headers          header.Headers `json:"headers,omitempty"`
	Cookies          header.Cookies `json:"cookies,omitempty"`
	ContentLength    int64          `json:"contentLength,omitempty"`
	ContentType      string         `json:"contentType,omitempty"`
	Charset          string         `json:"charset,omitempty"`
	TransferEncoding []string       `json:"transferEncoding,omitempty"`
	Close            bool           `json:"close,omitempty"`
	Uncompressed     bool           `json:"uncompressed,omitempty"`
	Body             *body.Body     `json:"body,omitempty"`
	Timing           *Timing        `json:"timing,omitempty"`
	Error            *string        `json:"error,omitempty"`
}

func ResponseFromHTTP(p *request.Prototype, r *http.Response, sess *session.Session, timing *Timing) *Response {
	headers := make(header.Headers, 0, len(r.Header))
	for k, vs := range r.Header {
		for _, v := range vs {
			headers = append(headers, &header.Header{Key: k, Value: v})
		}
	}
	headers.Sort()

	cookies := header.ParseCookies(r.Cookies())

	ct, charset := parseCT(headers.GetValue("Content-Type"))
	ce := headers.GetValue("Content-Encoding")
	bod, err := body.Parse(r.Request.URL.Path, ce, ct, charset, r.ContentLength, r.Body)
	var es *string = nil
	if err != nil {
		ex := err.Error()
		es = &ex
	}
	if len(ct) == 0 && bod != nil && bod.Config != nil {
		ct = bod.Config.MimeType()
	}
	return &Response{
		Method:           r.Request.Method,
		URL:              r.Request.URL.String(),
		RequestHeaders:   p.FinalHeaders(sess),
		Status:           r.Status,
		StatusCode:       r.StatusCode,
		Proto:            r.Proto,
		ProtoMajor:       r.ProtoMajor,
		ProtoMinor:       r.ProtoMinor,
		Headers:          headers,
		Cookies:          cookies,
		ContentLength:    r.ContentLength,
		ContentType:      ct,
		Charset:          charset,
		TransferEncoding: r.TransferEncoding,
		Close:            r.Close,
		Uncompressed:     r.Uncompressed,
		Body:             bod,
		Timing:           timing,
		Error:            es,
	}
}

func parseCT(h string) (string, string) {
	ct, cs := npncore.SplitString(h, ';', true)
	cs = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(strings.TrimSpace(cs), "charset"), "="))
	return strings.TrimSpace(ct), cs
}
