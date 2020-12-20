package call

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"logur.dev/logur"
)

type RequestStarted struct {
	Coll    string    `json:"coll"`
	Req     string    `json:"req"`
	ID      uuid.UUID `json:"id"`
	Idx     int       `json:"idx"`
	Method  string    `json:"method,omitempty"`
	URL     string    `json:"url,omitempty"`
	Started time.Time `json:"started,omitempty"`
}

type RequestCompleted struct {
	Coll     string    `json:"coll"`
	Req      string    `json:"req"`
	ID       uuid.UUID `json:"id"`
	Idx      int       `json:"idx"`
	Status   string    `json:"status,omitempty"`
	Rsp      *Response `json:"rsp,omitempty"`
	Error    string    `json:"error,omitempty"`
	Duration int       `json:"duration,omitempty"`
}

type Params struct {
	ID          uuid.UUID
	Idx         int
	UserID      *uuid.UUID
	Coll        string
	Req         string
	Client      *http.Client
	Proto       *request.Prototype
	Sess        *session.Session
	SessSvc     *session.Service
	Logger      logur.Logger
	OnStarted   func(started *RequestStarted)
	OnCompleted func(completed *RequestCompleted)
}

func (p *Params) Clone(proto *request.Prototype) *Params {
	return &Params{
		ID:          p.ID,
		Idx:         p.Idx + 1,
		UserID:      p.UserID,
		Coll:        p.Coll,
		Req:         p.Req,
		Client:      p.Client,
		Proto:       proto,
		Sess:        p.Sess,
		SessSvc:     p.SessSvc,
		Logger:      p.Logger,
		OnStarted:   p.OnStarted,
		OnCompleted: p.OnCompleted,
	}
}
