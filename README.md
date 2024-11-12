![workflow status](https://github.com/cswn/goCipher/actions/workflows/go.yml/badge.svg)

# goCipher

A fun command-line tool written in Go that encodes and decodes messages using various ciphers.

## Installation

To use the tool, you can either download the source code or use the built executable.

### Executable

Clone the repo: `git clone git@github.com:cswn/goCipher.git`

Enter the directory: `cd goCipher/`

Execute the tool: `./goCipher {subcommand} [ARGUMENTS]` -> see Usage

### Source

You must have [installed Go](https://go.dev/doc/install) locally to use the tool this way.

`git clone git@github.com:cswn/goCipher.git`

`cd goCipher/`

`go run main.go {subcommand} [ARGUMENTS]` -> see Usage

## Subcommands

### `ceasar`

The famous Ceasar cipher, one of the most well-known and simplest encryption techniques. Julius Ceasar is said to have used it himself for private communications. For this cipher, messages are encrypted by shifting the letters along the alphabet according to a numerical key.

Usage:

```shell
./goCipher ceasar -m "my secret message" -k 17
```

Arguments:

- `-m` The message to encrypt or decrypt. For strings with more than one word, make sure to enclose in quotes.
- `-d` (optional) Decrypt. Set this flag (with no arguments) if you want to decrypt instead of encrypt a message.
- `-k` (optional, default 10) The key, a number.

### `vigenere`

The classic Vigenère cipher, invented by Giovan Battista Bellaso in 1553 but mistakenly attributed to Blaise de Vigenère in the 16th century after he came up with a similar cipher. In this cipher, messages are encrypted using a secret key and a tabula recta. Because multiple substitution alphabets are used it is known as a 'polyalphabetic' cipher.

Usage:

```shell
./goCipher vigenere -m "my secret message" -k "lorem ipsum"
```

Arguments:

- `-m` The message to encrypt or decrypt. For strings with more than one word, make sure to enclose in quotes.
- `-k` The key, a string. For strings with more than one word, make sure to enclose in quotes.
- `-d` (optional) Decrypt. Set this flag (with no arguments) if you want to decrypt instead of encrypt a message.

### `playfair`

This cipher was the first to encrypt letters in pairs in cryptologic history. Its first recorded use was in a document from its inventor, Charles Wheatstone, in 1854. It uses a 5x5 matrix (or key table) containing a key word to encrypt or decrypt a message, digraph (pair of letters) by digraph.

Usage:

```shell
./goCipher playfair -m "attackatdawn" -k "lorem ipsum"
```

Arguments:

- `-m` The message to encrypt or decrypt. For strings with more than one word, make sure to enclose in quotes.
- `-k` The key, a string. For strings with more than one word, make sure to enclose in quotes.
- `-d` (optional) Decrypt. Set this flag (with no arguments) if you want to decrypt instead of encrypt a message.
