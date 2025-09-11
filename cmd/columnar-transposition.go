package cmd

import (
	"flag"
	"fmt"
	"math"
	"os"
	"unicode"

	"github.com/cswn/goCipher/internal"
)

type ColumnarSubCommand struct {
	message string
	decode  bool
	key     string
}

type Matrix [][]rune

func (cmd *ColumnarSubCommand) Name() string {
	return "columnar"
}

func (cmd *ColumnarSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The message to decode or encode")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
	flagSet.StringVar(&cmd.key, "k", "", "The key on which to decode and encode the message")
}

func (cmd *ColumnarSubCommand) Description() string {
	return "a transposition cipher, which re-arranges letters to form the ciphertext"
}

func (cmd *ColumnarSubCommand) Run() {
	if cmd.message == "" {
		fmt.Fprint(os.Stderr, "Please make sure to pass a message. \n")
		return
	}

	if cmd.key == "" {
		fmt.Fprint(os.Stderr, "Please make sure to pass a key. \n")
		return
	}

	newMsg := TransposeText(cmd.message, cmd.decode, cmd.key)
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, newMsg)
}

func TransposeText(plainText string, decode bool, key string) string {
	key = internal.FilterString(key, func(e rune) bool {
		return unicode.IsLetter(e) || unicode.IsSpace(e)
	})
	plainText = internal.FilterString(plainText, func(e rune) bool {
		return unicode.IsLetter(e) || unicode.IsSpace(e)
	})

	// create key table
	runes := []rune(plainText)

	return encryptTransposition(key, runes, decode)
}

func encryptTransposition(key string, msg []rune, decode bool) string {
	col := len(key)
	floatRows := float64(len(msg)) / float64(col)
	maxRows := int(math.Abs(floatRows))

	result := msg

	// pad end of runes
	fill := (maxRows * col) - len(msg)
	for i := 0; i < fill; i++ {
		result = append(result, 95)
	}

	return string(result)
}
