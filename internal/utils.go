package internal

import (
	"unicode"
)

const (
	CODE_POINT_A = 97
	CODE_POINT_I = 105
	CODE_POINT_X = 120
	CODE_POINT_Z = 122
)

// FilterString filters a string rune by rune using a callback function.
func FilterString(str string, callback func(rune) bool) string {
	var filtered = ""

	for _, item := range str {
		if callback(item) {
			filtered += string(item)
		}
	}

	return filtered
}

// RuneCountInStringLettersOnly counts the number of runes in a string that are letters.
func RuneCountInStringLettersOnly(text string) int {
	var totalRuneCount = 0

	for _, letter := range text {
		if unicode.IsLetter(letter) {
			totalRuneCount++
		}
	}

	return totalRuneCount
}
