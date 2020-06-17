package schematypes

import (
	"emperror.dev/errors"
	"encoding/json"
)

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
	case KeyBit:
		tgt := Bit{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyBool:
		tgt := Bool{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyByte:
		tgt := Byte{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyChar:
		tgt := Char{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyDate:
		tgt := Date{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyEnumValue:
		tgt := EnumValue{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyError:
		tgt := Error{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyFloat:
		tgt := Float{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyInt:
		tgt := Int{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyJSON:
		tgt := JSON{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyList:
		tgt := List{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyMap:
		tgt := Map{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyMethod:
		tgt := Method{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyOption:
		tgt := Option{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyReference:
		tgt := Reference{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeySet:
		tgt := Set{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyString:
		tgt := String{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTime:
		tgt := Time{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTimestamp:
		tgt := Timestamp{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTimestampZoned:
		tgt := TimestampZoned{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUnknown:
		tgt := Unknown{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUUID:
		tgt := UUID{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyXML:
		tgt := XML{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	default:
		t = Unknown{X: "unmarshal:" + wu.K}
	}
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal wrapped field of type [" + wu.K + "]")
	}
	if t == nil {
		return errors.New("nil type returned from unmarshal")
	}
	w.K = wu.K
	w.V = t
	return nil
}
