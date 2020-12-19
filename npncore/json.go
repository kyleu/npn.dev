package npncore

import (
	"bytes"
	"encoding/json"
	"fmt"

	"logur.dev/logur"
)

// Converts the argument to a string containing pretty JSON, logging errors
func ToJSON(x interface{}, logger logur.Logger) string {
	return string(ToJSONBytes(x, logger, true))
}

// Converts the argument to a string containing compact JSON, logging errors
func ToJSONCompact(x interface{}, logger logur.Logger) string {
	return string(ToJSONBytes(x, logger, false))
}

// Converts the argument to an optionally indented byte array, logging errors
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
		if logger != nil {
			logger.Warn(msg)
		}
	}
	return b
}

// Parses the provided JSON to the provided interface
func FromJSON(msg json.RawMessage, tgt interface{}) error {
	return json.Unmarshal(msg, tgt)
}

// Parses the provided JSON to the provided interface, validating that all fields are used
func FromJSONStrict(msg json.RawMessage, tgt interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(msg))
	dec.DisallowUnknownFields()
	return dec.Decode(tgt)
}

// Parses the provided JSON as a string
func FromJSONString(msg json.RawMessage) (string, error) {
	tgt := ""
	err := json.Unmarshal(msg, &tgt)
	if err != nil {
		return "", err
	}
	return tgt, nil
}
