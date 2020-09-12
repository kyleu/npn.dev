package call

import (
	"net/http"

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
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	ret := call(client, p, s.logger)

	return ret
}

func call(client *http.Client, p *request.Prototype, _ logur.Logger) *Result {
	req := p.ToHTTP()
	startNanos := npncore.StartTimer()
	hr, err := client.Do(req)
	duration := npncore.EndTimer(startNanos)

	status := "ok"
	var errStr *string = nil
	if err != nil {
		status = "error"
		es := err.Error()
		errStr = &es
	}

	rsp := ResponseFromHTTP(hr)

	return &Result{Status: status, Response: rsp, Duration: duration, Error: errStr}
}
