package schema

const (
	LevelInfo = iota
	LevelWarn
	LevelError
)

type ValidationMessage struct {
	Msg   string
	Level int
}

type ValidationResult struct {
	Schema   string
	Messages []ValidationMessage
	Duration int64
}

func validateSchema(s *Schema) ValidationResult {
	ret := ValidationResult{Schema: s.Key}
	e := func(msg string, level int) { ret.Messages = append(ret.Messages, ValidationMessage{Msg: msg, Level: level}) }
	e("Test error", LevelError)
	return ret
}
