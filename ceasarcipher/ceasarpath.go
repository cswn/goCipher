package ceasarcipher

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Path(input string) {
	// encode or decode
	var encodeChoice bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Would you like to encode or decode a message?")
	fmt.Println("[ e ] encode")
	fmt.Println("[ d ] decode")

	encodeChoiceInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	encodeChoiceInput = strings.TrimSuffix(encodeChoiceInput, "\n")
	if encodeChoiceInput == "d" {
		encodeChoice = false
	} else if encodeChoiceInput == "e" {
		encodeChoice = true
	} else {
		fmt.Println("Invalid choice")
		return
	}

	// key
	fmt.Print("Type your key (number) > ")
	keyReader := bufio.NewReader(os.Stdin)
	stringKey, err := keyReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	stringKey = strings.TrimSuffix(stringKey, "\n")
	intKey, err := strconv.Atoi(stringKey)
	if err != nil {
		panic(errors.New("that wasn't a number"))
	}

	// message
	fmt.Print("Type your message > ")
	plainTextReader := bufio.NewReader(os.Stdin)
	plainText, err := plainTextReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	plainText = strings.TrimSuffix(plainText, "\n")

	if !encodeChoice {
		Decode(plainText, intKey)
	} else if encodeChoice {
		Encode(plainText, intKey)
	}
}
