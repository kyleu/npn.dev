package schematypes

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
)

type Type interface {
	Key() string
	fmt.Stringer
}

type Wrapped struct {
	K string `json:"k"`
	V Type   `json:"t,omitempty"`
}

func Wrap(t Type) Wrapped {
	_, ok := t.(Wrapped)
	if ok {
		return t.(Wrapped)
	}
	return Wrapped{K: t.Key(), V: t}
}

func (w Wrapped) Key() string {
	return w.K
}

func (w Wrapped) String() string {
	return w.V.String()
}

type wrappedUnmarshal struct {
	K string `json:"k"`
	T json.RawMessage `json:"t"`
}

func (w *Wrapped) UnmarshalJSON(data []byte) error {
	var wu wrappedUnmarshal
	err := json.Unmarshal(data, &wu)
	if err != nil {
		return err
	}
	var t Type
	switch wu.K {
	case KeyBool:
		tgt := Bool{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyInt:
		tgt := Int{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyFloat:
		tgt := Float{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyString:
		tgt := String{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyMethod:
		tgt := Method{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyList:
		tgt := List{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUnion:
		tgt := Union{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyReference:
		tgt := Reference{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyOption:
		tgt := Option{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyError:
		tgt := Error{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUnknown:
		tgt := Unknown{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	default:
		t = Unknown{X: "unmarshal:" + wu.K}
	}
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal")
	}
	if t == nil {
		return errors.New("nil type returned from unmarshal")
	}
	w.K = wu.K
	w.V = t
	return nil
}
