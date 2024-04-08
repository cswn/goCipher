package main

import (
	"fmt"
	"os"

	goCipher "github.com/cswn/goCipher/cmd"
)

func main() {
	err := goCipher.Main()
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
