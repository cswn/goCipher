![workflow status](https://github.com/cswn/ciphertool/actions/workflows/go.yml/badge.svg)

# Cipher Tool

A fun command-line tool written in Go that encodes and decodes messages using various ciphers.

## Installation

To use the tool, you can either download the source code or use the built executable.

### Executable

Clone the repo: `git clone git@github.com:cswn/cipherTool.git`

Enter the directory: `cd cipherTool/`

Execute the tool: `./cipherTool {subcommand} [ARGUMENTS]` -> see Usage

### Source

You must have [installed Go](https://go.dev/doc/install) locally to use the tool this way.

`git clone git@github.com:cswn/cipherTool.git`

`cd cipherTool/`

`go run main.go {subcommand} [ARGUMENTS]` -> see Usage

## Subcommands

### `ceasar`

The famous Ceasar cipher, one of the most well-known and simplest encryption techniques. Julius Ceasar is said to have used it himself for private communications. For this cipher, messages are encrypted by shifting the letters along the alphabet according to a numerical key.

Usage:

```shell
./cipherTool ceasar -m "my secret message" -k 17
```

Arguments:

- `-m` The message to encrypt or decrypt. For strings with more than one word, make sure to enclose in quotes.
- `-d` (optional) Decrypt. Set this flag (with no arguments) if you want to decrypt instead of encrypt a message.
- `-k` (optional, default 10) The key, a number.

### `vigenere`

The classic Vigenère cipher, invented by Giovan Battista Bellaso in 1553 but mistakenly attributed to Blaise de Vigenère in the 16th century after he came up with a similar cipher. In this cipher, messages are encrypted using a secret key and a tabula recta. Because multiple substitution alphabets are used it is known as a 'polyalphabetic' cipher.

Usage:

```shell
./cipherTool vigenere -m "my secret message" -k "lorem ipsum"
```

Arguments:

- `-m` The message to encrypt or decrypt. For strings with more than one word, make sure to enclose in quotes.
- `-k` The key, a string. For strings with more than one word, make sure to enclose in quotes.
- `-d` (optional) Decrypt. Set this flag (with no arguments) if you want to decrypt instead of encrypt a message.
