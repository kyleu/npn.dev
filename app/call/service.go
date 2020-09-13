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
		e := "no request"
		return &Result{
			Status: "error",
			Error:  &e,
		}
	}
	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
		Timeout:   60 * time.Second,
	}

	ret := call(client, p, s.logger)

	return ret
}

func call(client *http.Client, p *request.Prototype, _ logur.Logger) *Result {
	req := p.ToHTTP()
	timing := &Timing{}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), timing.Trace()))
	timing.Begin()
	hr, err := client.Do(req)
	timing.Complete()

	status := "ok"
	var errStr *string = nil
	if err != nil {
		status = "error"
		es := err.Error()
		errStr = &es
	}

	rsp := ResponseFromHTTP(hr)

	return &Result{Status: status, Response: rsp, Timing: timing, Error: errStr}
}
