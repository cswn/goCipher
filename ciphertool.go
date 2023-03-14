package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/claraswan/cipherTool/ceasarcipher"
	"github.com/claraswan/cipherTool/vigenerecipher"
)

func main() {
	fmt.Println("Welcome to the ciper tool.")
	fmt.Print("Type c to use Ceaser cipher or v to use Vigenère cipher: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occurred")
		return
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.ToLower(input)

	if input == "c" {
		ceasarcipher.Path(input)
	} else if input == "v" {
		vigenerecipher.Path(input)
	} else {
		fmt.Println("Response not accepted.")
		return
	}
}
