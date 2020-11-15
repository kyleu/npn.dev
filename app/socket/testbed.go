package socket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
)

type tkv struct {
	T string          `json:"t"`
	K string          `json:"k"`
	V json.RawMessage `json:"v"`
}

func testbed(input json.RawMessage) error {
	parsed := &tkv{}
	err := npncore.FromJSON(input, parsed)
	if err != nil {
		return errors.Wrap(err, "cannot parse input")
	}
	return process(parsed)
}

func process(x *tkv) error {
	switch x.T {
	case "theme":
		return onTheme(x.K, x.V)
	default:
		return errors.New("unhandled testbed type [" + x.T + "]")
	}
}

func onTheme(k string, v json.RawMessage) error {
	var s string
	err := npncore.FromJSON(v, &s)
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
