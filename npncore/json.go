package npncore

import (
	"bytes"
	"encoding/json"
	"fmt"

	"logur.dev/logur"
)

func ToJSON(x interface{}, logger logur.Logger) string {
	return string(ToJSONBytes(x, logger, true))
}

func ToJSONCompact(x interface{}, logger logur.Logger) string {
	return string(ToJSONBytes(x, logger, false))
}

func ToJSONBytes(x interface{}, logger logur.Logger, indent bool) []byte {
	var b []byte
	var err error
	if indent {
		b, err = json.MarshalIndent(x, "", "  ")
	} else {
		b, err = json.Marshal(x)
	}
	if err != nil {
		msg := fmt.Sprintf("unable to serialize json from type [%T]: %+v", x, err)
		if logger == nil {
			println(msg)
		} else {
			logger.Warn(msg)
		}
	}
	return b
}

func FromJSON(msg json.RawMessage, tgt interface{}) error {
	return json.Unmarshal(msg, tgt)
}

func FromJSONStrict(msg json.RawMessage, tgt interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(msg))
	dec.DisallowUnknownFields()
	return dec.Decode(tgt)
}
