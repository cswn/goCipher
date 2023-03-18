package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/claraswan/cipherTool/ceasarcipher"
	"github.com/claraswan/cipherTool/vigenerecipher"
)

func main() {
	fmt.Println("Welcome to the ciper tool.")
	fmt.Println("Choose which cipher you'd like to use:")
	fmt.Println("[ c ] Ceasar cipher")
	fmt.Println("[ v ] Vigenere cipher")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(errors.New("a read error occurred"))
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.ToLower(input)

	if input == "c" {
		ceasarcipher.Path()
	} else if input == "v" {
		vigenerecipher.Path()
	} else {
		fmt.Println("Response not accepted.")
		return
	}
}
