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

func (w *Body) UnmarshalJSON(data []byte) error {
	x := &bodyJSON{}
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	if x == nil {
		return nil
	}
	w.Type = x.Type
	switch w.Type {
	case KeyTemp:
		temp := &Temp{}
		err = json.Unmarshal(x.Config, &temp)
		if err != nil {
			return err
		}
		w.Config = temp
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
