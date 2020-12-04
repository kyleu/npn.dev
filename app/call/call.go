package call

import (
	"fmt"
	"net/http/httptrace"
	"time"

	"golang.org/x/text/language"

	"github.com/kyleu/npn/npncore"
)

func call(p *CallParams) *Result {
	httpReq := p.Proto.ToHTTP(p.Sess)
	timing := &Timing{}
	httpReq = httpReq.WithContext(httptrace.WithClientTrace(httpReq.Context(), timing.Trace()))
	url := httpReq.URL.String()

	reqStartedMsg := &RequestStarted{Coll: p.Coll, Req: p.Req, ID: p.ID, Idx: p.Idx, Method: p.Proto.Method.Key, URL: url, Started: time.Now()}
	p.OnStarted(reqStartedMsg)

	p.Logger.Info("making call to [" + url + "]")
	timing.Begin()

	hr, err := p.Client.Do(httpReq)

	status := "ok"
	var errStr = ""
	if err != nil {
		status = "error"
		errStr = err.Error()
	}

	timing.CompleteHeaders()

	var rsp *Response
	if hr != nil {
		rsp = ResponseFromHTTP(p.Proto, hr, p.Sess, timing)
		parseCookies := p.Proto.Options == nil || (!p.Proto.Options.IgnoreCookies)
		if parseCookies && p.Sess != nil && len(rsp.Cookies) > 0 {
			if p.Sess.AddCookies(rsp.Cookies...) {
				if p.SessSvc != nil {
					// p.Logger.Debug(fmt.Sprintf("saving session [%v] (%v cookies, %v variables)", p.Sess.Key, len(p.Sess.Cookies), len(p.Sess.Variables)))
					err = p.SessSvc.Save(p.UserID, p.Sess.Key, p.Sess)
					if err != nil {
						p.Logger.Warn(fmt.Sprintf("unable to save session: %+v", err))
					}
				}
			}
		}
	}

	timing.Complete()

	ret := NewResult(p.ID, p.Coll, p.Req, status)
	ret.Response = rsp
	ret.Error = errStr

	p.Logger.Info("call to [" + url + "] complete in [" + npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed) + "]")

	reqCompleteMsg := &RequestCompleted{Coll: p.Coll, Req: p.Req, ID: p.ID, Idx: p.Idx, Status: status, Rsp: ret.Response, Error: errStr, Duration: timing.Completed}
	p.OnCompleted(reqCompleteMsg)

	shouldRedir := p.Proto.Options == nil || (!p.Proto.Options.IgnoreRedirects)
	redir := rsp != nil && rsp.StatusCode >= 300 && rsp.StatusCode < 400 && rsp.Headers.Contains("location")

	if shouldRedir && redir {
		redirP := getRedir(rsp, p.Proto)
		if redirP == nil {
			return ret
		}
		p.Logger.Debug("redirecting to [" + redirP.URLString() + "]")
		redirParams := p.Clone(redirP)
		return call(redirParams)
	}

	return ret
}
