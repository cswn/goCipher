package cmd

import "unicode"

func filterString(str string, callback func(rune) bool) string {
	var filtered = ""

	for _, item := range str {
		if callback(item) {
			filtered += string(item)
		}
	}

	return filtered
}

func runeCountInStringLettersOnly(text string) int {
	totalRuneCount := 0
	for _, letter := range text {
		if unicode.IsLetter(letter) {
			totalRuneCount++
		}
	}
	return totalRuneCount
}
