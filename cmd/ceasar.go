package cmd

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

type CeasarSubCommand struct {
	message string
	decode  bool
	key     int64
}

func (cmd *CeasarSubCommand) Name() string {
	return "ceasar"
}

func (cmd *CeasarSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The message to decode or encode")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
	flagSet.Int64Var(&cmd.key, "k", 10, "The key on which to decode and encode the message")
}

func (cmd *CeasarSubCommand) Description() string {
	return "the famous ceasar cipher, where messages are decoded by shifting the letters along the alphabet according to a key"
}

func (cmd *CeasarSubCommand) Run(args []string) {
	// todo if no message was passed, return an error and print usage
	newMsg := ShiftText(cmd.message, cmd.decode, cmd.key)
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, newMsg)
}

func ShiftText(plainText string, decode bool, shiftKey int64) string {
	if decode {
		shiftKey = -shiftKey
	}

	cipherText := ""
	plainText = strings.ToLower(plainText)

	for _, letter := range plainText {
		if !unicode.IsLetter(letter) {
			cipherText += string(letter)
		} else {
			var newPos int64

			if (int64(letter) + shiftKey) > 122 {
				newPos = (97 + (shiftKey % 26))
			} else if (int64(letter) + shiftKey) < 97 {
				newPos = (122 + (shiftKey % 26))
			} else {
				newPos = (int64(letter) + (shiftKey % 26))
			}

			cipherText += string(rune(newPos))
		}
	}

	return cipherText
}
