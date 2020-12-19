package npncore

import (
	"regexp"
	"strings"
)

var regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-._]")
var regexpMultipleDashes = regexp.MustCompile("-+")

// Converts a string to a URL-safe representation, replacing forbidden charaters with "-"
func Slugify(s string) (slug string) {
	slug = strings.TrimSpace(s)

	slug = strings.ToLower(slug)

	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, "-")
	slug = regexpMultipleDashes.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-.")

	slug = smartTruncate(slug)

	return slug
}

func smartTruncate(text string) string {
	maxLength := 256
	if len(text) < maxLength {
		return text
	}

	var truncated string
	words := strings.SplitAfter(text, "-")
	if len(words[0]) > maxLength {
		return words[0][:maxLength]
	}

	for _, word := range words {
		if len(truncated)+len(word)-1 <= maxLength {
			truncated += word
		} else {
			break
		}
	}
	return strings.Trim(truncated, "-")
}
