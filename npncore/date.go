package npncore

import (
	"time"

	"emperror.dev/errors"
)

// Year, month, day
const YMD = "2006-01-02"

// Full date format
const DateFull = "2006-01-02 15:04:05"

// Const set to 24
const HoursInDay = 24

// Returns a string representing this time.Time in YMD format
func ToYMD(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(YMD)
}

// Parses a string represention of a time.Time in YMD format
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

// Parses a string represention of a time.Time in DateFull format
func ToDateString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(DateFull)
}
