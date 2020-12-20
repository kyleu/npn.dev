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

// Stack frame definition
type ErrorFrame struct {
	Key string
	Loc string
}

type ErrorFrames []*ErrorFrame

// An error's message, stack, and cause
type ErrorDetail struct {
	Message    string
	StackTrace errors.StackTrace
	Cause      *ErrorDetail
}

// Creates an ErrorDetail for the provided error
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

// Converts a stack trace to a set of ErrorFrames
func TraceDetail(trace errors.StackTrace) ErrorFrames {
	s := fmt.Sprintf("%+v", trace)
	lines := strings.Split(s, "\n")
	validLines := make([]string, 0)

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 {
			validLines = append(validLines, l)
		}
	}

	ret := make(ErrorFrames, 0)

	for i := 0; i < len(validLines)-1; i += 2 {
		f := &ErrorFrame{Key: validLines[i], Loc: validLines[i+1]}
		ret = append(ret, f)
	}

	return ret
}

// Makes error messages for a provided ID string
func IDErrorString(k string, v string) string {
	if len(v) == 0 {
		return fmt.Sprintf("empty %v id", k)
	}
	return fmt.Sprintf("invalid %v id [%v]", k, v)
}

// Helper for making errors related to an ID
func IDError(k string, v string) error {
	return errors.New(IDErrorString(k, v))
}
