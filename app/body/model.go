package body

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"io"
	"io/ioutil"
	"logur.dev/logur"
)

type Config interface {
	ContentLength() int64
	Bytes() []byte
	MimeType() string
	String() string
	Merge(data npncore.Data, logger logur.Logger) Config
}

type Body struct {
	Type   string `json:"type"`
	Length int64  `json:"length"`
	Config Config `json:"config"`
}

func NewBody(t string, c Config) *Body {
	return &Body{Type: t, Length: c.ContentLength(), Config: c}
}

func (b *Body) String() string {
	s := "nil"
	if b.Config != nil {
		s = b.Config.String()
	}
	return fmt.Sprintf("%v:%v", b.Type, s)
}

type bodyJSON struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

func (b *Body) UnmarshalJSON(data []byte) error {
	x := &bodyJSON{}
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	if x == nil {
		return nil
	}
	b.Type = x.Type
	switch b.Type {
	case KeyError:
		e := &Error{}
		err = json.Unmarshal(x.Config, &e)
		if err != nil {
			return err
		}
		b.Config = e
	case KeyForm:
		f := &Form{}
		err = json.Unmarshal(x.Config, &f)
		if err != nil {
			return err
		}
		b.Config = f
	case KeyHTML:
		h := &HTML{}
		err = json.Unmarshal(x.Config, &h)
		if err != nil {
			return err
		}
		b.Config = h
	case KeyImage:
		i := &Image{}
		err = json.Unmarshal(x.Config, &i)
		if err != nil {
			return err
		}
		b.Config = i
	case KeyJSON:
		js := &JSON{}
		err = json.Unmarshal(x.Config, &js)
		if err != nil {
			return err
		}
		b.Config = js
	case KeyLarge:
		l := &Large{}
		err = json.Unmarshal(x.Config, &l)
		if err != nil {
			return err
		}
		b.Config = l
	case KeyRaw:
		raw := &Raw{}
		err = json.Unmarshal(x.Config, &raw)
		if err != nil {
			return err
		}
		b.Config = raw
	case "":
		return nil
	default:
		return errors.New("invalid body type [" + x.Type + "]")
	}
	return nil
}

func (b *Body) ToHTTP() io.ReadCloser {
	if b == nil || b.Config == nil {
		return nil
	}
	return ioutil.NopCloser(bytes.NewReader(b.Config.Bytes()))
}

func (b *Body) ContentLength() int64 {
	if b == nil || b.Config == nil {
		return 0
	}
	return b.Config.ContentLength()
}

func (b *Body) Merge(data npncore.Data, logger logur.Logger) *Body {
	if b == nil || len(b.Type) == 0 {
		return nil
	}
	cfg := b.Config
	if cfg != nil {
		cfg = b.Config.Merge(data, logger)
	}
	return &Body{
		Type: npncore.MergeLog("body.type", b.Type, data, logger),
		Length: b.Length,
		Config: cfg,
	}
}
