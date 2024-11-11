package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/TwiN/go-color"
)

type PlayfairSubCommand struct {
	message string
	decode  bool
	key     string
}

type KeyTable [][]rune

func (cmd *PlayfairSubCommand) Name() string {
	return "playfair"
}

func (cmd *PlayfairSubCommand) Flags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cmd.message, "m", "", "The message to decode or encode")
	flagSet.BoolVar(&cmd.decode, "d", false, "Decode the message instead of encoding")
	flagSet.StringVar(&cmd.key, "k", "", "The key on which to decode and encode the message")
}

func (cmd *PlayfairSubCommand) Description() string {
	return "a polygraphic substitution cipher, which encrypts digraphs instead of single letters"
}

func (cmd *PlayfairSubCommand) Run() {
	if cmd.message == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a message. \n"))
		return
	}

	if cmd.key == "" {
		fmt.Fprint(os.Stderr, color.Ize(color.Red, "Please make sure to pass a key. \n"))
		return
	}

	newMsg := ShiftTextByDigraph(cmd.message, cmd.decode, cmd.key)
	encodedOrDecoded := "encoded"
	if cmd.decode {
		encodedOrDecoded = "decoded"
	}
	fmt.Printf("Your %s message is: %s \n", encodedOrDecoded, newMsg)
}

func ShiftTextByDigraph(plainText string, decode bool, key string) string {
	// remove whitespaces and non-letter runes from the key
	key = strings.ReplaceAll(key, " ", "")
	key = FilterString(key, func(e rune) bool {
		return unicode.IsLetter(e)
	})

	// create key table
	keyTable := generateKeyTable(key, len(key))
	fmt.Println("keyTable: %v", keyTable)

	preparedRunes := preparePlainText(plainText)
	encrypted := encrypt(preparedRunes, keyTable, len(preparedRunes))
	//fmt.Println("encrypted: %v", encrypted)

	return encrypted
}

func generateKeyTable(key string, keyLength int) KeyTable {
	kt := make(KeyTable, 5)

	// make all the rows of table
	for row := range kt {
		kt[row] = make([]rune, row)
	}

	// fill with null values
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if len(kt[j]) == 5 {
				continue
			}
			kt[i] = append(kt[j], 0)
		}
	}

	// add letters of the key to keytable first
	runesLeftInKey := keyLength
	for i := 0; i < keyLength; i++ {
		if runesLeftInKey != 0 {
			for j := 0; j < keyLength; j++ {
				letter := rune(key[j])
				kt[i][j] = letter
				runesLeftInKey = runesLeftInKey - 1
			}
		}
	}

	// then fill the rest of the table with the letters of alphabet in order
	maxCodePoint := CODE_POINT_A + (26 - (keyLength % 26))

	// iterate through each row of table
	for i := 0; i < 5; i++ {
		rowIndex := 0

		// iterate through letters a-z
		for j := CODE_POINT_A; j < maxCodePoint; j++ {
			letter := rune(j)
			exists := runeExistsInMap(letter, kt)
			if rowIndex == 5 {
				continue
			}

			// fmt.Println("letter", letter)
			// fmt.Println("rowIndex", rowIndex)
			// fmt.Println("kt[i][rowIndex]", kt[i][rowIndex])

			// if the letter doesn't already exist in the current row, and there is
			// no value set yet, set the value to be the letter and go to next position in row
			if !exists && kt[i][rowIndex] == 0 {
				kt[i][rowIndex] = letter
				rowIndex++
			}
		}
	}

	return kt
}

func preparePlainText(plainText string) []rune {
	if len(plainText)%2 != 0 {
		plainText += "z"
	}

	arr := make([]rune, len(plainText))

	for i, letter := range plainText {
		arr[i] = letter
	}

	return arr
}

func searchForDigraphInKeyTable(kt KeyTable, a rune, b rune) [4]int {
	var result [4]int
	var i int
	var j int

	if a == 'j' {
		a = CODE_POINT_I
	} else if b == 'j' {
		b = CODE_POINT_I
	}

	// go through rows
	for i = 0; i < 5; i++ {
		// go through columns
		for j = 0; j < 5; j++ {
			if kt[i][j] == a {
				result[0] = i
				result[1] = j
			} else if kt[i][j] == b {
				result[2] = i
				result[3] = j
			}
		}
	}

	return result
}

func encrypt(msg []rune, kt KeyTable, messageLength int) string {
	for i := 0; i < messageLength; i += 2 {
		digraphRectangle := searchForDigraphInKeyTable(kt, msg[i], msg[i+1])
		fmt.Println("digraphRectangle:", digraphRectangle)
	}

	return "hi"
}

// check if a letter is already present in the keytable
func runeExistsInMap(letter rune, table KeyTable) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if table[i][j] == letter {
				return true
			}
		}
	}

	return false
}
