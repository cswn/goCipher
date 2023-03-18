package ceasarcipher

import (
	"fmt"
	"strings"
	"unicode"
)

func ShiftText(plainText string, shift int) string {
	cipherText := ""
	plainText = strings.ToLower(plainText)
	// Still doesn't work when you choose a key that is under 26, but
	// still goes past rune num 122 (which is z)

	for _, letter := range plainText {
		if !unicode.IsLetter(letter) {
			cipherText += string(letter)
		} else {
			newutf8 := (int(letter) + (shift % 26))
			cipherText += string(rune(newutf8))
		}
	}

	return cipherText
}

func Encode(plainText string, shift int) {
	encodedMsg := ShiftText(plainText, shift)
	fmt.Printf("Your encoded message is: %s \n", encodedMsg)
	ContinuePath()
}

func Decode(plainText string, shift int) {
	decodedMsg := ShiftText(plainText, -shift)
	fmt.Printf("Your decoded message is: %s \n", decodedMsg)
	ContinuePath()
}
