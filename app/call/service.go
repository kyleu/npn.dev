package call

import (
	"net/http"
	"net/http/httptrace"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/language"

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
		return NewErrorResult(coll, req, "no request")
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
		rsp = ResponseFromHTTP(p, hr, timing)
		rsp.Prior = prior
	}
	if rsp == nil {
		rsp = prior
	}

	timing.Complete()

	ret := NewResult(coll, req, status)
	ret.Response = rsp
	ret.Error = errStr

	logger.Info("call to [" + url + "] complete in [" + npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed) + "]")

	if p.Options == nil || (!p.Options.IgnoreRedirects) {
		if rsp != nil && rsp.StatusCode >= 300 && rsp.StatusCode < 400 && rsp.Headers.Contains("location") {
			loc := rsp.Headers.GetValue("location")
			if len(loc) == 0 {

			}
			if strings.HasPrefix(loc, "//") {
				loc = p.Protocol.Key + ":" + loc
			}
			if !strings.Contains(loc, "://") {
				if !strings.HasPrefix(loc, "/") {
					loc = filepath.Dir(p.Path) + "/" + loc
				}
				loc = p.Protocol.Key + "://" + p.Host() + loc
			}
			redirP := request.PrototypeFromString(loc)
			redirP.Auth = p.Auth
			redirP.Headers = p.Headers
			redirP.Options = p.Options
			logger.Debug("redirecting to [" + redirP.URLString() + "]")
			return call(coll, req, client, redirP, ret.Response, logger)
		}
	}

	return ret
}
