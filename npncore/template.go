package npncore

import (
	"emperror.dev/errors"
	"fmt"
	"logur.dev/logur"
	"strings"
)

var defaultPrefix = "{"
var defaultSuffix = "}"

func Merge(content string, args Data) (string, error) {
	return mergeVariables(content, args, defaultPrefix, defaultSuffix, 0)
}

func MergeNeeded(key string) bool {
	return strings.Contains(key, defaultPrefix)
}

func MergeLog(key string, content string, args Data, logger logur.Logger) string {
	x, err := Merge(content, args)
	if err != nil {
		logger.Warn(fmt.Sprintf("unable to merge [%v] %+v", key, err))
		return content
	}
	return x
}

func mergeVariables(content string, args Data, start string, end string, depth int) (string, error) {
	if depth > 32 {
		return content, errors.New("template recursion error for [" + content + "]")
	}
	sIdx := strings.Index(content, start)
	if sIdx > -1 {
		eIdx := strings.Index(content[sIdx:], end)
		if eIdx > -1 {
			orig := content[sIdx : sIdx+eIdx+len(end)]

			n := orig[len(start) : len(orig)-len(end)]
			d := ""
			dIdx := strings.Index(orig, "|")
			if dIdx > -1 {
				n = orig[len(start):dIdx]
				d = orig[dIdx+1 : len(orig)-len(end)]
			}

			o := args.GetString(n)
			if len(o) == 0 || o == "<nil>" {
				o = d
			}
			if len(o) == 0 || o == "<nil>" {
				o = n
			}
			return mergeVariables(strings.Replace(content, orig, o, 1), args, start, end, depth+1)
		}
	}

	return content, nil
}

var tests = []struct {
	Src  string
	Data Data
	Tgt  string
}{
	{Src: "a{b}c", Data: nil, Tgt: "abc"},                                  // Missing
	{Src: "a{b}c", Data: Data{"b": "x"}, Tgt: "axc"},                       // Basic
	{Src: "a{b}c", Data: Data{"b": "{foo}zz", "foo": "xx"}, Tgt: "axxzzc"}, // Recursive
	{Src: "a{b|default}c", Data: nil, Tgt: "adefaultc"},                    // Default
	{Src: "a{b|default}c", Data: Data{"b": "x"}, Tgt: "axc"},               // Skip default
}

func MergeTests() error {
	for _, t := range tests {
		r, err := Merge(t.Src, t.Data)
		if err != nil {
			return errors.Wrap(err, "merge error for ["+t.Src+"]")
		}
		if r != t.Tgt {
			return errors.New(fmt.Sprintf("merge expected [%v] but observed [%v]", t.Tgt, r))
		}
	}
	return nil
}
