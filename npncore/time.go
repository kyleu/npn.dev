package npncore

import (
	"time"
)

func StartTimer() int64 {
	return time.Now().UnixNano()
}

func EndTimer(startNanos int64) int {
	return int((time.Now().UnixNano() - startNanos) / int64(time.Microsecond))
}
