package body

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type Config interface {
	ContentLength() int64
	Bytes() []byte
	MimeType() string
	String() string
}

type Body struct {
	Type   string `json:"type"`
	Config Config `json:"config"`
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
	if b == nil {
		return nil
	}
	return ioutil.NopCloser(bytes.NewReader(b.Config.Bytes()))
}

func (b *Body) ContentLength() int64 {
	if b == nil {
		return 0
	}
	return b.Config.ContentLength()
}
