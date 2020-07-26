package npncore

import (
	"time"
)

const YMD = "2006-01-02"
const DateFull = "2006-01-02 15:04:05"

func ToYMD(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(YMD)
}

func ToDateString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(DateFull)
}
