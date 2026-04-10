package challenge6

import (
	"strings"
	"unicode"
)

func CountWordFrequency(text string) map[string]int {
	result := make(map[string]int)

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			return r
		}
		if r == '-' {
			return ' '
		}
		return -1
	}, text)

	for _, word := range strings.Fields(cleaned) {
		result[strings.ToLower(word)]++
	}

	return result
}