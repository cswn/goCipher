package ceasarcipher

import (
	"fmt"
	"strings"
	"unicode"
)

func ShiftText(plainText string, shift int) {
	cipherText := ""
	plainText = strings.ToLower(plainText)

	for _, letter := range plainText {
		//fmt.Println(int(letter))
		if !unicode.IsLetter(letter) {
			cipherText += string(letter)
		} else {
			newutf8 := (int(letter) + (shift % 26))
			fmt.Println(rune(newutf8))
			cipherText += string(rune(newutf8))
		}
	}

	fmt.Println(cipherText)
}

func Encode(plainText string, shift int) {
	ShiftText(plainText, shift)
}

func Decode(plainText string, shift int) {
	ShiftText(plainText, -shift)
}
