package ceasarcipher

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Path() {
	// encode or decode
	var encodeChoice bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Would you like to encode or decode a message?")
	fmt.Println("[ e ] encode")
	fmt.Println("[ d ] decode")

	encodeChoiceInput, err := reader.ReadString('\n')
	if err != nil {
		panic(errors.New("an error occured while reading your input"))
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
		panic(errors.New("an error occured while reading your input"))
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
		panic(errors.New("an error occured while reading your input"))
	}
	plainText = strings.TrimSuffix(plainText, "\n")

	if !encodeChoice {
		Decode(plainText, intKey)
	} else if encodeChoice {
		Encode(plainText, intKey)
	}
}

func ContinuePath() {
	fmt.Print("Do you want to en/decode another message? y/n > ")
	continueReader := bufio.NewReader(os.Stdin)
	continu, err := continueReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured")
		return
	}
	continu = strings.TrimSuffix(continu, "\n")
	continu = strings.ToLower(continu)

	if continu == "y" {
		Path()
	} else if continu == "n" {
		fmt.Println("Ok, bye!")
		return
	} else {
		ContinuePath()
		fmt.Println("Please choose either 'y' or 'n'")
		return
		// call continuePath again?
	}

}
