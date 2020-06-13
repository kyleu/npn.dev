package util

import (
	"encoding/json"
	"fmt"

	"logur.dev/logur"
)

func ToJSON(x interface{}, logger logur.Logger) string {
	return string(ToJSONBytes(x, logger))
}

func ToJSONBytes(x interface{}, logger logur.Logger) []byte {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil && logger != nil {
		logger.Warn(fmt.Sprintf("unable to serialize json from type [%T]: %+v", x, err))
	}
	return b
}

func FromJSON(msg json.RawMessage, tgt interface{}, logger logur.Logger) {
	err := json.Unmarshal(msg, tgt)
	if err != nil && logger != nil {
		logger.Warn(fmt.Sprintf("error unmarshalling JSON [%v]: %+v", string(msg), err))
	}
}
