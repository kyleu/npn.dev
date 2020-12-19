package npncore

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Converts 10,000,000 to "10,000ms"
func MicrosToMillis(l language.Tag, i int) string {
	div := 1000
	min := 20

	ms := i / div
	if ms >= min {
		return FormatInteger(l, ms) + "ms"
	}

	x := float64(ms) + (float64(i%div) / float64(div))
	p := message.NewPrinter(l)

	return p.Sprintf("%.3f", x) + "ms"
}

// Formats an integer with the provided language tag
func FormatInteger(l language.Tag, v int) string {
	p := message.NewPrinter(l)
	return p.Sprintf("%d", v)
}
