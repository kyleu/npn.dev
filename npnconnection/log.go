package npnconnection

// Message intended to be sent to Connection objects representing a log message
type LogMessage struct {
	Level   string            `json:"level"`
	Message string            `json:"message"`
	Context map[string]string `json:"context,omitempty"`
}

// Constructor
func NewLogMessage(level string, msg string, ctx ...string) *LogMessage {
	c := make(map[string]string, len(ctx)/2)
	for i := 0; i < len(ctx); i += 2 {
		c[ctx[i]] = ctx[i+1]
	}
	return &LogMessage{Level: level, Message: msg, Context: c}
}
