package main

import (
	"bufio"
	"fmt"
	"github.com/claraswan/cipherTool/ceasarcipher"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the ciper tool.")
	fmt.Print("Type c to use ceaser cipher: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.ToLower(input)
	if input != "c" {
		fmt.Println("Response not accepted.")
		return
	}
	ceasarcipher.ceasarCipherPath(input)
}
