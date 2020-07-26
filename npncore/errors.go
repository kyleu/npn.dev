package npncore

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type unwrappable interface {
	Unwrap() error
}

type ErrorFrame struct {
	Key string
	Loc string
}

type ErrorDetail struct {
	Message    string
	StackTrace errors.StackTrace
	Cause      *ErrorDetail
}

func GetErrorDetail(e error) *ErrorDetail {
	var stack errors.StackTrace = nil

	t, ok := e.(stackTracer)
	if ok {
		stack = t.StackTrace()
	}

	var cause *ErrorDetail = nil

	u, ok := e.(unwrappable)
	if ok {
		cause = GetErrorDetail(u.Unwrap())
	}

	return &ErrorDetail{
		Message:    e.Error(),
		StackTrace: stack,
		Cause:      cause,
	}
}

func TraceDetail(trace errors.StackTrace) []ErrorFrame {
	s := fmt.Sprintf("%+v", trace)
	lines := strings.Split(s, "\n")
	validLines := make([]string, 0)

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 {
			validLines = append(validLines, l)
		}
	}

	ret := make([]ErrorFrame, 0)

	for i := 0; i < len(validLines)-1; i += 2 {
		f := ErrorFrame{Key: validLines[i], Loc: validLines[i+1]}
		ret = append(ret, f)
	}

	return ret
}

func IDErrorString(k string, v string) string {
	if len(v) == 0 {
		return fmt.Sprintf("invalid %v id", k)
	}
	return fmt.Sprintf("invalid %v id \"%v\"", k, v)
}

func IDError(k string, v string) error {
	return errors.New(IDErrorString(k, v))
}
