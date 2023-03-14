package ceasarcipher

import (
	"fmt"
	"strings"
)

func CeasarCipher(msg string, key string, encode bool) string {
	if encode {
		fmt.Println("encoding: ", msg)

		encodedMsg := ""
		msg = strings.ToLower(msg)
		return encodedMsg
	} else {
		fmt.Println("decoding: ", msg)

		decodedMsg := ""
		return decodedMsg
	}
}
