/*
GoCipher en- and decrypts messages using simple, common ciphers over the command line.
It is meant for personal and for-fun use only.
Usage:

	goCipher [subcommand] [arguments]

Subcommands are:

	ceasar
		The famous ceasar cipher, a simple alphabetic shift cipher.
	vigenere
		The classic Vigenère cipher, a polyalphabetic cipher that used a key and a tabula recta.
	playfair
		A polygraphic substitution cipher, which encrypts digraphs instead of single letters.

If you choose to run the subcommand without its optional arguments, the default values will be taken.
*/
package main

import (
	"fmt"
	"os"

	goCipher "github.com/cswn/goCipher/cmd"
)

func main() {
	err := goCipher.Main()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
