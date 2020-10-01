package call

import (
	"golang.org/x/text/language"
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

func (s *Service) Call(coll string, req string, p *request.Prototype) *Result {
	if p == nil {
		return NewErrorResult(p.URLString(), coll, req, "no request")
	}
	return call(coll, req, getClient(p), p, nil, s.logger)
}

func getClient(p *request.Prototype) *http.Client {
	timeout := 60 * time.Second
	if p.Options != nil && p.Options.Timeout > 0 {
		timeout = time.Duration(p.Options.Timeout) * time.Second
	}
	return &http.Client{
		Transport: &http.Transport{},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: timeout,
	}
}

func call(coll string, req string, client *http.Client, p *request.Prototype, prior *Response, logger logur.Logger) *Result {
	httpReq := p.ToHTTP()
	timing := &Timing{}
	httpReq = httpReq.WithContext(httptrace.WithClientTrace(httpReq.Context(), timing.Trace()))
	url := httpReq.URL.String()
	logger.Info("making call to [" + url + "]")
	timing.Begin()

	hr, err := client.Do(httpReq)

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
		rsp.Prior = prior
	}
	if rsp == nil {
		rsp = prior
	}

	timing.Complete()

	ret := NewResult(url, coll, req, status)
	ret.RequestHeaders = p.FinalHeaders()
	ret.Response = rsp
	ret.Timing = timing
	ret.Error = errStr

	logger.Info("call to [" + url + "] complete in [" + npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed) + "]")

	// TODO handle redirects, recurse with prior
	return ret
}
