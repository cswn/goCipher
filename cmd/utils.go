package cmd

import (
	"unicode"
)

const (
	CODE_POINT_A = 97
	CODE_POINT_Z = 122
)

func FilterString(str string, callback func(rune) bool) string {
	var filtered = ""

	for _, item := range str {
		if callback(item) {
			filtered += string(item)
		}
	}

	return filtered
}

func RuneCountInStringLettersOnly(text string) int {
	totalRuneCount := 0
	for _, letter := range text {
		if unicode.IsLetter(letter) {
			totalRuneCount++
		}
	}
	return totalRuneCount
}
