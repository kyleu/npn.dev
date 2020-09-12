package body

import (
	"encoding/json"
	"errors"
	"fmt"
)

type BodyConfig interface {
	Bytes() []byte
	String() string
}

type Body struct {
	Type   string     `json:"type"`
	Config BodyConfig `json:"config"`
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
	w.Type = x.Type
	switch w.Type {
	case KeyTemp:
		temp := &Temp{}
		err = json.Unmarshal(x.Config, &temp)
		if err != nil {
			return err
		}
		w.Config = temp
	default:
		return errors.New("invalid body type [" + x.Type + "]")
	}
	return nil
}
