package cmd

import (
	"flag"
)

type ColumnarSubCommand struct {
	message string
	decode  bool
	key     string
}

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

func (cmd *ColumnarSubCommand) Run() {}
