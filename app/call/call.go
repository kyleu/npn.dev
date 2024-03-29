package call

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"

	"golang.org/x/text/language"

	"github.com/kyleu/libnpn/npncore"
)

func call(p *Params) error {
	httpReq := p.Proto.ToHTTP(p.Sess)
	timing := &Timing{}
	httpReq = httpReq.WithContext(httptrace.WithClientTrace(httpReq.Context(), timing.Trace()))
	u := httpReq.URL.String()

	reqStartedMsg := &RequestStarted{Coll: p.Coll, Req: p.Req, ID: p.ID, Idx: p.Idx, Method: p.Proto.Method.Key, URL: u, Started: time.Now()}
	p.OnStarted(reqStartedMsg)

	p.Logger.Info("making call to [" + u + "]")
	timing.Begin()

	hr, err := p.Client.Do(httpReq)
	if err != nil {
		timing.Complete()
		msg := err.Error()
		if strings.Contains(msg, "Client.Timeout exceeded while awaiting headers") {
			msg = fmt.Sprintf("timeout after [%v]", npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed))
		}
		reqCompleteMsg := &RequestCompleted{Coll: p.Coll, Req: p.Req, ID: p.ID, Idx: p.Idx, Status: "ok", Error: msg, Duration: timing.Completed}
		p.OnCompleted(reqCompleteMsg)
		return err
	}

	timing.CompleteHeaders()

	var rsp *Response
	if hr != nil {
		rsp = parseResponse(hr, p, timing)
	}

	timing.Complete()

	p.Logger.Info("call to [" + u + "] complete in [" + npncore.MicrosToMillis(language.AmericanEnglish, timing.Completed) + "]")

	reqCompleteMsg := &RequestCompleted{Coll: p.Coll, Req: p.Req, ID: p.ID, Idx: p.Idx, Status: "ok", Rsp: rsp, Duration: timing.Completed}
	p.OnCompleted(reqCompleteMsg)

	shouldRedir := p.Proto.Options == nil || (!p.Proto.Options.IgnoreRedirects)
	redir := rsp != nil && rsp.StatusCode >= 300 && rsp.StatusCode < 400 && rsp.Headers.Contains("location")

	if shouldRedir && redir {
		redirP := getRedir(rsp, p.Proto)
		if redirP == nil {
			return nil
		}
		p.Logger.Debug("redirecting to [" + redirP.URLString() + "]")
		redirParams := p.Clone(redirP)
		return call(redirParams)
	}

	return nil
}

func parseResponse(hr *http.Response, p *Params, timing *Timing) *Response {
	rsp := ResponseFromHTTP(p.Proto, hr, p.Sess, timing)
	parseCookies := p.Proto.Options == nil || (!p.Proto.Options.IgnoreCookies)
	if parseCookies && p.Sess != nil && len(rsp.Cookies) > 0 {
		addCookies(p, rsp)
	}
	return rsp
}

func addCookies(p *Params, rsp *Response) {
	if p.Sess.AddCookies(rsp.Cookies...) {
		if p.SessSvc != nil {
			// p.Logger.Debug(fmt.Sprintf("saving session [%v] (%v cookies, %v variables)", p.Sess.Key, len(p.Sess.Cookies), len(p.Sess.Variables)))
			err := p.SessSvc.Save(p.UserID, p.Sess.Key, p.Sess)
			if err != nil {
				p.Logger.Warn(fmt.Sprintf("unable to save session: %+v", err))
			}
		}
	}
}
