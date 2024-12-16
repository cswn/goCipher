package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/cswn/goCipher/internal"
)

type SingleByteXorSubCommand struct {
	message string
	decode  bool
	key     uint
}

func (cmd *SingleByteXorSubCommand) Name() string {
	return "single byte xor"
}

func (cmd *SingleByteXorSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The hexidecimal string to decode or encode")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
	flagSet.UintVar(&cmd.key, "k", 0, "The byte on which to decode and encode the message")
}

func (cmd *SingleByteXorSubCommand) Description() string {
	return "XOR (exclusive or) is a logical operation used as an encryption method in modern cryptography"
}

func (cmd *SingleByteXorSubCommand) Run() {
	if cmd.message == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a message. \n"))
		return
	}

	bytes := []byte(cmd.message)
	var err error
	if cmd.decode {
		bytes, err = internal.DecodeHex([]byte(cmd.message))
		if err != nil {
			fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure your message is in hexadecimal. \n"))
			return
		}
	}
	plainTextBytes := []byte(bytes)

	if cmd.key == 0 {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a key. \n"))
		return
	}

	if cmd.key > 256 {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make your key is no more than one byte long. \n"))
		return
	}

	new := XORCipher(plainTextBytes, byte(cmd.key))
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message in hexadecimal is: %s \n", encodedOrDecoded, internal.EncodeHex(new))
}

func XORCipher(bytes []byte, key byte) []byte {
	result := make([]byte, len(bytes))
	for i := range len(bytes) {
		result[i] = bytes[i] ^ key
	}
	return result
}
