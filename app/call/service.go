package call

import (
	"github.com/gofrs/uuid"
	"net/http"
	"time"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	logger logur.Logger
	sessSvc *session.Service
}

func NewService(sessSvc *session.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: "call"})
	return &Service{sessSvc: sessSvc, logger: logger}
}

func (s *Service) Call(userID *uuid.UUID, coll string, req string, p *request.Prototype, sess *session.Session, started func(started *RequestStarted), completed func(completed *RequestCompleted)) *Result {
	rID := npncore.UUID()
	if p == nil {
		return NewErrorResult(rID, coll, req, "no request")
	}
	return call(&CallParams{
		ID:      rID,
		UserID:  userID,
		Coll:    coll,
		Req:     req,
		Client:  getClient(p),
		Proto:   p,
		Sess:    sess,
		SessSvc: s.sessSvc,
		Logger:  s.logger,
		OnStarted: started,
		OnCompleted: completed,
	})
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
