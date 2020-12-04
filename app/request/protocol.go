package request

import (
	"encoding/json"
	"strings"
)

type Protocol struct {
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
}

var (
	ProtocolHTTP  = Protocol{Key: "http", Description: ""}
	ProtocolHTTPS = Protocol{Key: "https", Description: ""}
	ProtocolWS    = Protocol{Key: "ws", Description: ""}
	ProtocolWSS   = Protocol{Key: "wss", Description: ""}
)
var AllProtocols = []Protocol{ProtocolHTTP, ProtocolHTTPS, ProtocolWS, ProtocolWSS}

func ProtocolFromString(s string) Protocol {
	for _, t := range AllProtocols {
		if t.Key == s {
			return t
		}
	}
	return Protocol{Key: s, Description: "Custom protocol"}
}

func (t Protocol) String() string {
	return t.Key
}

func (t Protocol) Secure() bool {
	return strings.HasSuffix(t.Key, "s")
}

func (t *Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Protocol) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = ProtocolFromString(s)
	return nil
}
