package vigenerecipher

import "fmt"

func VigenereCipher(msg string, key string, encode bool) {
	if encode == true {
		fmt.Print("encoding", msg, key)
	} else {
		fmt.Print("decoding", msg, key)
	}
}
