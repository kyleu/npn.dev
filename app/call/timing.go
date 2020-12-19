package call

import (
	"crypto/tls"
	"net/http/httptrace"
	"time"

	"github.com/kyleu/npn/npncore"
)

type TimingSection struct {
	Key   string `json:"key"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type Timing struct {
	Began             int64 `json:"began,omitempty"`
	DNSStart          int   `json:"dnsStart,omitempty"`
	DNSEnd            int   `json:"dnsEnd,omitempty"`
	ConnectStart      int   `json:"connectStart,omitempty"`
	ConnectEnd        int   `json:"connectEnd,omitempty"`
	TLSStart          int   `json:"tlsStart,omitempty"`
	TLSEnd            int   `json:"tlsEnd,omitempty"`
	WroteHeaders      int   `json:"wroteHeaders,omitempty"`
	WroteRequest      int   `json:"wroteRequest,omitempty"`
	FirstResponseByte int   `json:"firstResponseByte,omitempty"`
	ResponseHeaders   int   `json:"responseHeaders,omitempty"`
	Completed         int   `json:"completed,omitempty"`
}

func (t *Timing) Begin() {
	t.Began = time.Now().UnixNano()
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

func (t *Timing) ConnectComplete() int {
	if t.TLSEnd > 0 {
		return t.TLSEnd
	}
	return t.ConnectEnd
}

func (t *Timing) CompleteHeaders() {
	t.ResponseHeaders = int((npncore.TimerStart() - t.Began) / 1000)
}

func (t *Timing) Complete() {
	t.Completed = int((npncore.TimerStart() - t.Began) / 1000)
}

func (t *Timing) Duration() int {
	return t.Completed
}

func (t *Timing) Trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			t.DNSStart = int((npncore.TimerStart() - t.Began) / 1000)
		},
		DNSDone: func(httptrace.DNSDoneInfo) {
			t.DNSEnd = int((npncore.TimerStart() - t.Began) / 1000)
		},
		ConnectStart: func(string, string) {
			t.ConnectStart = int((npncore.TimerStart() - t.Began) / 1000)
		},
		ConnectDone: func(string, string, error) {
			t.ConnectEnd = int((npncore.TimerStart() - t.Began) / 1000)
		},
		TLSHandshakeStart: func() {
			t.TLSStart = int((npncore.TimerStart() - t.Began) / 1000)
		},
		TLSHandshakeDone: func(tls.ConnectionState, error) {
			t.TLSEnd = int((npncore.TimerStart() - t.Began) / 1000)
		},
		WroteHeaders: func() {
			t.WroteHeaders = int((npncore.TimerStart() - t.Began) / 1000)
		},
		WroteRequest: func(httptrace.WroteRequestInfo) {
			t.WroteRequest = int((npncore.TimerStart() - t.Began) / 1000)
		},
		GotFirstResponseByte: func() {
			t.FirstResponseByte = int((npncore.TimerStart() - t.Began) / 1000)
		},
	}
}

func (t *Timing) Sections() []*TimingSection {
	ret := make([]*TimingSection, 0, 5)
	var add = func(k string, s int, e int) {
		ret = append(ret, &TimingSection{Key: k, Start: s, End: e})
	}
	add("dns", t.DNSStart, t.DNSEnd)
	add("connect", t.ConnectStart, t.ConnectEnd)
	if t.TLSEnd > 0 {
		add("tls", t.TLSStart, t.TLSEnd)
	}
	add("reqheaders", t.ConnectComplete(), t.WroteHeaders)
	add("reqbody", t.WroteHeaders, t.WroteRequest)
	add("rspwait", t.WroteRequest, t.FirstResponseByte)
	add("rspheaders", t.FirstResponseByte, t.ResponseHeaders)
	add("rspbody", t.ResponseHeaders, t.Completed)
	return ret
}
