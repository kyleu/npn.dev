package call

import (
	"net/http"
	"time"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
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

func (s *Service) Call(coll string, req string, p *request.Prototype, sess *session.Session) *Result {
	if p == nil {
		return NewErrorResult(coll, req, "no request")
	}
	return call(coll, req, getClient(p), p, nil, sess, s.logger)
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
