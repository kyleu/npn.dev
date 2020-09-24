package call

import (
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	logger logur.Logger
}

func NewService(logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: "call"})
	return &Service{logger: logger}
}

func (s *Service) Call(p *request.Prototype) *Result {
	if p == nil {
		return &Result{Status: "error", Error: "no request"}
	}
	return call(getClient(p), p, s.logger)
}

func getClient(p *request.Prototype) *http.Client {
	timeout := 60 * time.Second
	if p.Options != nil && p.Options.Timeout > 0 {
		timeout = time.Duration(p.Options.Timeout) * time.Second
	}
	return &http.Client{
		Transport: &http.Transport{},
		Timeout:   timeout,
	}
}

func call(client *http.Client, p *request.Prototype, _ logur.Logger) *Result {
	req := p.ToHTTP()
	timing := &Timing{}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), timing.Trace()))
	timing.Begin()
	hr, err := client.Do(req)

	status := "ok"
	var errStr string = ""
	if err != nil {
		status = "error"
		errStr = err.Error()
	}

	timing.CompleteHeaders()

	var rsp *Response
	if hr != nil {
		rsp = ResponseFromHTTP(hr)
	}

	timing.Complete()

	return &Result{Status: status, Response: rsp, Timing: timing, Error: errStr}
}
