package npncore

import (
	"time"
)

// Returns current time as nanos
func TimerStart() int64 {
	return time.Now().UnixNano()
}

// Returns the difference between the provided start time and current time in nanoseconds, as microseconds
func TimerEnd(startNanos int64) int {
	return int((time.Now().UnixNano() - startNanos) / int64(time.Microsecond))
}
