package call

import (
	"context"
	"net"
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	logger  logur.Logger
	sessSvc *session.Service
}

func NewService(sessSvc *session.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: "call"})
	return &Service{sessSvc: sessSvc, logger: logger}
}

func (s *Service) Call(userID *uuid.UUID, coll string, req string, p *request.Prototype, sess *session.Session, started func(started *RequestStarted), completed func(completed *RequestCompleted)) error {
	rID := npncore.UUID()
	if p == nil {
		return errors.New("nil request passed to [call]")
	}
	return call(&Params{
		ID:          rID,
		UserID:      userID,
		Coll:        coll,
		Req:         req,
		Client:      getClient(p),
		Proto:       p.Merge(sess.Variables.ToData(), s.logger),
		Sess:        sess,
		SessSvc:     s.sessSvc,
		Logger:      s.logger,
		OnStarted:   started,
		OnCompleted: completed,
	})
}

var d = &net.Dialer{Timeout: 4 * time.Second}
var transport = &http.Transport{
	DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return d.Dial(network, addr)
	},
	DialTLSContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return d.Dial(network, addr)
	},
	TLSHandshakeTimeout:   2 * time.Second,
	IdleConnTimeout:       1 * time.Second,
	ResponseHeaderTimeout: 0,
	ExpectContinueTimeout: 0,
}

func getClient(p *request.Prototype) *http.Client {
	timeout := 10 * time.Second
	if p.Options != nil && p.Options.Timeout > 0 {
		timeout = time.Duration(p.Options.Timeout) * time.Second
	}

	return &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: timeout,
	}
}
