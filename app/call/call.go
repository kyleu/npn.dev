package call

import (
	"github.com/kyleu/npn/app/session"
	"golang.org/x/text/language"
	"net/http"
	"net/http/httptrace"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

func call(coll string, req string, client *http.Client, p *request.Prototype, prior *Response, sess *session.Session, logger logur.Logger) *Result {
	httpReq := p.ToHTTP()
	timing := &Timing{}
	httpReq = httpReq.WithContext(httptrace.WithClientTrace(httpReq.Context(), timing.Trace()))
	url := httpReq.URL.String()
	logger.Info("making call to [" + url + "]")
	timing.Begin()

	hr, err := client.Do(httpReq)

	status := "ok"
	var errStr = ""
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

	sess.AddCookies(rsp.Cookies)

	logger.Info("call to [" + url + "] complete in [" + npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed) + "]")

	shouldRedir := p.Options == nil || (!p.Options.IgnoreRedirects)
	redir := rsp != nil && rsp.StatusCode >= 300 && rsp.StatusCode < 400 && rsp.Headers.Contains("location")

	if shouldRedir && redir {
		redirP := getRedir(rsp, p)
		if redirP == nil {
			return ret
		}
		logger.Debug("redirecting to [" + redirP.URLString() + "]")
		return call(coll, req, client, redirP, ret.Response, sess, logger)
	}

	return ret
}
