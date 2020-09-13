package call

import (
	"crypto/tls"
	"github.com/kyleu/npn/npncore"
	"net/http/httptrace"
	"time"
)

type Timing struct {
	began             int64
	DNSStart          int `json:"dnsStart,omitempty"`
	DNSEnd            int `json:"dnsEnd,omitempty"`
	ConnectStart      int `json:"connectStart,omitempty"`
	ConnectEnd        int `json:"connectEnd,omitempty"`
	TLSStart          int `json:"tlsStart,omitempty"`
	TLSEnd            int `json:"tlsEnd,omitempty"`
	WroteHeaders      int `json:"wroteHeaders,omitempty"`
	WroteRequest      int `json:"wroteRequest,omitempty"`
	FirstResponseByte int `json:"firstResponseByte,omitempty"`
	Completed         int `json:"complete,omitempty"`
}

func (t *Timing) Begin() {
	t.began = time.Now().UnixNano()
}

func (t *Timing) DNS() (int, int) {
	return t.DNSStart, t.DNSEnd - t.DNSStart
}

func (t *Timing) TLS() (int, int) {
	return t.TLSStart, t.TLSEnd - t.TLSStart
}

func (t *Timing) Connect() (int, int) {
	return t.ConnectStart, t.ConnectEnd - t.ConnectStart
}

func (t *Timing) Request() (int, int, int) {
	return t.FirstResponseByte, t.WroteHeaders - t.FirstResponseByte, t.WroteRequest - t.WroteHeaders
}

func (t *Timing) Duration() int {
	return t.Completed
}

func (t *Timing) Trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			t.DNSStart = int((npncore.StartTimer() - t.began) / 1000)
		},
		DNSDone: func(httptrace.DNSDoneInfo) {
			t.DNSEnd = int((npncore.StartTimer() - t.began) / 1000)
		},
		ConnectStart: func(string, string) {
			t.ConnectStart = int((npncore.StartTimer() - t.began) / 1000)
		},
		ConnectDone: func(string, string, error) {
			t.ConnectEnd = int((npncore.StartTimer() - t.began) / 1000)
		},
		TLSHandshakeStart: func() {
			t.TLSStart = int((npncore.StartTimer() - t.began) / 1000)
		},
		TLSHandshakeDone: func(tls.ConnectionState, error) {
			t.TLSEnd = int((npncore.StartTimer() - t.began) / 1000)
		},
		WroteHeaders: func() {
			t.WroteHeaders = int((npncore.StartTimer() - t.began) / 1000)
		},
		WroteRequest: func(httptrace.WroteRequestInfo) {
			t.WroteRequest = int((npncore.StartTimer() - t.began) / 1000)
		},
		GotFirstResponseByte: func() {
			t.FirstResponseByte = int((npncore.StartTimer() - t.began) / 1000)
		},
	}
}

func (t *Timing) Complete() {
	t.Completed = int((npncore.StartTimer() - t.began) / 1000)
}
