package npncore

import (
	"errors"
	"strings"
)

func Template(content string, args Data) (string, error) {
	return parseTemplate(content, args, "{{", "}}", 0)
}

func SplitString(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

func SplitStringLast(s string, sep byte, cutc bool) (string, string) {
	i := strings.LastIndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

func parseTemplate(content string, args Data, start string, end string, depth int) (string, error) {
	if depth > 32 {
		return content, errors.New("template recursion error for [" + content + "]")
	}
	sIdx := strings.Index(content, start)
	if sIdx > -1 {
		eIdx := strings.Index(content[sIdx:], end)
		if eIdx > -1 {
			orig := content[sIdx:sIdx + eIdx + len(end)]

			n := orig[len(start):len(orig) - len(end)]
			d := ""
			dIdx := strings.Index(orig, "|")
			if dIdx > -1 {
				n = orig[len(start):dIdx]
				d = orig[dIdx + 1:len(orig) - 1]
			}

			o := args.GetString(n)
			if len(o) == 0 {
				o = d
			}
			return parseTemplate(strings.Replace(content, orig, o, 1), args, start, end, depth + 1)
		}
	}

	return content, nil
}
