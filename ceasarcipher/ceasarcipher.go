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
	// still goes past rune 122 (which is z)
	// need to define range of runes that it should stay within (97 -122)

	for _, letter := range plainText {
		if !unicode.IsLetter(letter) {
			cipherText += string(letter)
		} else {
			var newPos int
			if (int(letter) + shift) > 122 {
				newPos = (97 + (shift % 26))
			} else if (int(letter) + shift) < 97 {
				newPos = (122 + (shift % 26))
			} else {
				newPos = (int(letter) + (shift % 26))
			}

			cipherText += string(rune(newPos))
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
