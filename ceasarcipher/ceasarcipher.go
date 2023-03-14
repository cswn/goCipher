package ceasarcipher

import (
	"fmt"
	"strings"
	"unicode"
)

func CeasarCipher(msg string, key int64, encode bool) {
	var alphabet = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	if encode {
		encodedMsg := ""
		msg = strings.ToLower(msg)

		for _, letter := range msg {
			if unicode.IsLetter(letter) {
				indexInAlphabet := indexOf(string(letter), alphabet[:])
				bigIndex := (int64(indexInAlphabet) + key) % 26
				newLetter := alphabet[bigIndex]
				encodedMsg += newLetter
			} else {
				encodedMsg += string(letter)
			}
		}

		fmt.Println(encodedMsg)

	} else {
		fmt.Println("decoding: ", msg)

		decodedMsg := ""
		fmt.Println(decodedMsg)
	}
}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
