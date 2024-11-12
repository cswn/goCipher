package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/TwiN/go-color"
	"github.com/cswn/goCipher/internal"
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
	return "the famous ceasar cipher, where messages are encrypted by shifting the letters along the alphabet according to a key"
}

func (cmd *CeasarSubCommand) Run() {
	if cmd.message == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a message. \n"))
		return
	}

	newMsg := ShiftText(cmd.message, cmd.decode, cmd.key)
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, newMsg)
}

// ShiftText takes plaintext and "shifts" it across the 26-letter alphabet, either to the right or to the left, with
// number of steps shifted according to the shiftKey.
func ShiftText(plainText string, decode bool, shiftKey int64) string {
	if decode {
		shiftKey = -shiftKey
	}

	cipherText := ""
	plainText = strings.ToLower(plainText)

	// a 'for range' loop decodes one UTF-8-encoded rune on each iteration
	for _, letter := range plainText {
		if !unicode.IsLetter(letter) {
			cipherText += string(letter)
		} else {
			// normalize the key to shift by w/ modulo 26
			shiftKey = shiftKey % 26
			var newCodePoint int64 = int64(letter) + shiftKey

			// if the letter's position added to the key goes past the code point for 'z' or 'a',
			// capture the difference between them and redirect it to the beginning or end of alphabet
			if newCodePoint > internal.CODE_POINT_Z {
				newCodePoint = internal.CODE_POINT_A + (newCodePoint - internal.CODE_POINT_Z) - 1
			} else if newCodePoint < internal.CODE_POINT_A {
				newCodePoint = internal.CODE_POINT_Z - (internal.CODE_POINT_A - newCodePoint) + 1
			}

			cipherText += string(rune(newCodePoint))
		}
	}

	return cipherText
}
