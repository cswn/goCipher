package cmd

import (
	"flag"
	"fmt"
	"os"

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

func (cmd *VigenereSubCommand) Run(args []string) {
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

func ShiftTextByKeyword(message string, decode bool, key string) string {

}
