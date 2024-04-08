package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

type SubCommand interface {
	Name() string
	Flags(*flag.FlagSet)
	Description() string
	Run()
}

var subcommands map[string]SubCommand = make(map[string]SubCommand)

func init() {
	subcommands["ceasar"] = &CeasarSubCommand{}
	subcommands["vigenere"] = &VigenereSubCommand{}
}

func usageForSubCommand(subcommand SubCommand) string {
	usage := "    " + (color.InBlue(subcommand.Name()))
	usage += fmt.Sprintf("\n      %s\n", subcommand.Description())
	return usage
}

func usage() string {
	usage := "goCipher is a command line tool for encoding and decoding fun secret messages.\n"
	usage += "Available commands:\n"
	for _, subcommand := range subcommands {
		usage += usageForSubCommand(subcommand)
	}
	return usage
}

func Main() error {
	args := os.Args

	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		return errors.New(color.InRed("You must provide a valid subcommand.\n"))
	}

	subcommand := args[1]

	if subcommand == "help" {
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		return nil
	}

	cmd, ok := subcommands[subcommand]
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		return errors.New(color.InRed("You must provide a valid subcommand.\n"))
	}

	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.Flags(flagSet)
	err := flagSet.Parse(args[2:])
	if err != nil {
		return errors.New(color.InRed("There was an error: " + err.Error() + "\n"))
	}

	cmd.Run()

	return nil
}
