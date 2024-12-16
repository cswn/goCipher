package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

type SingleByteXorSubCommand struct {
	message string
	decode  bool
	key     string
}

func (cmd *SingleByteXorSubCommand) Name() string {
	return "single byte xor"
}

func (cmd *SingleByteXorSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The message to decode or encode")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
	flagSet.StringVar(&cmd.key, "k", "", "The byte on which to decode and encode the message")
}

func (cmd *SingleByteXorSubCommand) Description() string {
	return "XOR (exclusive or) is a logical operation used as an encryption method in modern cryptography"
}

func (cmd *SingleByteXorSubCommand) Run() {
	keyBytes := []byte(cmd.key)
	plainTextBytes := []byte(cmd.message)

	if cmd.message == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a message. \n"))
		return
	}

	if len(cmd.key) == 0 {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a key. \n"))
		return
	}

	if len(keyBytes) > 1 {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make your key is only one byte long. \n"))
		return
	}

	new := XORCipher(plainTextBytes, keyBytes[0])
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, string(new))
}

func XORCipher(bytes []byte, key byte) []byte {
	result := make([]byte, len(bytes))
	for i := range len(bytes) {
		result[i] = bytes[i] ^ key
	}
	return result
}
