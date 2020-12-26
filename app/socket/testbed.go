package socket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kyleu/libnpn/npnconnection"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npncore"
)

type tkv struct {
	T string          `json:"t"`
	K string          `json:"k"`
	V json.RawMessage `json:"v"`
}

func testbed(input json.RawMessage, connSvc *npnconnection.Service) error {
	parsed := &tkv{}
	err := npncore.FromJSON(input, parsed)
	if err != nil {
		return errors.Wrap(err, "cannot parse input")
	}
	return process(parsed, connSvc)
}

func process(x *tkv, connSvc *npnconnection.Service) error {
	switch x.T {
	case "theme":
		return onTheme(x.K, x.V)
	case "log":
		return onLog(x.K, x.V, connSvc)
	default:
		return errors.New("unhandled testbed type [" + x.T + "]")
	}
}

func onLog(level string, json json.RawMessage, svc *npnconnection.Service) error {
	content, err := npncore.FromJSONString(json)
	if err != nil {
		return errors.Wrap(err, "invalid content")
	}

	err = svc.BroadcastLog(level, content)
	if err != nil {
		return errors.Wrap(err, "error broadcasting log")
	}
	return nil
}

func onTheme(k string, v json.RawMessage) error {
	s, err := npncore.FromJSONString(v)
	if err != nil {
		return errors.Wrap(err, "unable to process theme value")
	}

	fn := "./ui/src/user/themes.ts"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return errors.Wrap(err, "unable to read ["+fn+"]")
	}

	content := string(b)

	startIdx := strings.Index(content, k)
	newContent := ""
	if startIdx == -1 {
		startIdx = strings.LastIndex(content, "export const allThemes")
		endIdx := startIdx
		post := content[endIdx:]
		pi := strings.Index(post, "\n];")
		post = post[0:pi] + ", " + k + post[pi:]
		newContent = content[0:(startIdx)] + s + ";\n\n" + post
	} else {
		endIdx := strings.Index(content[startIdx+1:], ";") + startIdx
		newContent = content[0:(startIdx-6)] + s + content[endIdx+1:]
	}

	println(fmt.Sprintf("saving themes with [%v] bytes", len(newContent)))

	// println(newContent)
	err = ioutil.WriteFile(fn, []byte(newContent), 0)

	return err
}
