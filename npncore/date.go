package npncore

import (
	"emperror.dev/errors"
	"time"
)

const YMD = "2006-01-02"
const DateFull = "2006-01-02 15:04:05"
const HoursInDay = 24

func ToYMD(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(YMD)
}

func FromYMD(s string) (*time.Time, error) {
	if len(s) == 0 {
		return nil, nil
	}
	ret, err := time.Parse(YMD, s)
	if err != nil {
		return nil, errors.New("invalid date [" + s + "] (expected 2020-01-15)")
	}
	return &ret, nil
}

func ToDateString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(DateFull)
}
