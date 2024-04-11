package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/TwiN/go-color"
)

type VigenereSubCommand struct {
	message string
	decode  bool
	key     string
}

func (cmd *VigenereSubCommand) Name() string {
	return "vigenere"
}

func (cmd *VigenereSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The message to decode or encode")
	flagSet.StringVar(&cmd.key, "k", "", "The key on which to decode and encode the message")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
}

func (cmd *VigenereSubCommand) Description() string {
	return "the classic Vigenère cipher, where messages are encrypted using a secret key and a tabula recta"
}

func (cmd *VigenereSubCommand) Run() {
	// todo if no message was passed, return an error and print usage
	if cmd.message == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a message. \n"))
		return
	}

	if cmd.key == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a key. \n"))
		return
	}

	newMsg := ShiftTextByKeyword(cmd.message, cmd.decode, cmd.key)
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, newMsg)
}

// ShiftTextByKeyword takes plaintext and "shifts" each letter separately, according to the position in the alphabet
// of the letter in the key which has the same index as the current plaintext letter.
// To ensure the key is long enough to provide this for each letter of the plaintext, it is padded by repeating itself until it
// is at least the length of the plaintext string.
func ShiftTextByKeyword(plainText string, decode bool, key string) string {
	var runeCountPlainText = RuneCountInStringLettersOnly(plainText)

	// remove whitespaces and non-letter runes from the key
	key = strings.ReplaceAll(key, " ", "")
	key = FilterString(key, func(e rune) bool {
		return unicode.IsLetter(e)
	})

	// pad the the key string so that it's longer than plaintext
	var lengthAdjustedKey string
	for runeCountPlainText > utf8.RuneCountInString(lengthAdjustedKey) {
		lengthAdjustedKey += key
	}

	// call ShiftText except "shiftKey" is the rune number of each letter
	var message string
	var adjustedIndex int = 0
	for _, letter := range plainText {
		letterKey := lengthAdjustedKey[adjustedIndex] - byte(CODE_POINT_A)

		if !unicode.IsLetter(letter) { // make sure the next letter in the key is used and not skipped over in case of a non-letter
			message += string(letter)
			adjustedIndex = adjustedIndex - 1
		} else {
			message += ShiftText(string(letter), decode, int64(letterKey))
		}

		adjustedIndex++
	}

	return message
}
