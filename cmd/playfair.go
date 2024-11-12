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
	// remove whitespaces and non-letter runes from the key and plainText
	key = strings.ReplaceAll(key, " ", "")
	plainText = strings.ReplaceAll(plainText, " ", "")
	key = FilterString(key, func(e rune) bool {
		return unicode.IsLetter(e)
	})
	plainText = FilterString(plainText, func(e rune) bool {
		return unicode.IsLetter(e)
	})

	// create key table
	keyTable := generateKeyTable(key, len(key))
	fmt.Println(keyTable)

	// prepare/pad the plaintext so it's able to be extracted into pairs of letters
	preparedRunes := preparePlainText(plainText)

	// encrypt
	encrypted := encrypt(preparedRunes, keyTable, len(preparedRunes))

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
	keyIndex := 0
	for i := 0; i < keyLength; i++ {
		if runesLeftInKey == 0 {
			break
		}

		j := 0
		for runesLeftInKey > 0 {
			if j == 5 {
				break
			}

			letter := rune(key[keyIndex])
			exists := runeExistsInMap(letter, kt)

			if !exists {
				kt[i][j] = letter
				runesLeftInKey = runesLeftInKey - 1
				keyIndex++
			}
			j++
		}
	}

	// then fill the rest of the table with the letters of alphabet in order
	// iterate through each row of table
	for i := 0; i < 5; i++ {
		rowIndex := 0

		// for each row of the table, iterate through letters a-z
		// max number of iterations should be amount of letters remaining between
		j := CODE_POINT_A
		for j < CODE_POINT_Z+1 {
			// don't go over 5 items per row
			if rowIndex == 5 {
				break
			}
			// if there is already a rune in place, continue
			if kt[i][rowIndex] != 0 {
				rowIndex++
				continue
			}

			// the letter j (code point 106) should not be included in the key table
			if j == 106 {
				j++
				continue
			}

			// extract letter and check if it is already in the keytable
			letter := rune(j)
			exists := runeExistsInMap(letter, kt)

			// if the letter doesn't already exist in the current row, and there is
			// no value set yet, set the value to be the letter and go to next position in row
			if !exists {
				kt[i][rowIndex] = letter
				rowIndex++
			}

			j++
		}
	}

	return kt
}

func preparePlainText(plainText string) []rune {
	// make sure no two same letters would be in the same digraph
	// if they are, pad them with a bogus letter, x
	// first make sure the loop will not go out of range
	max := len(plainText)
	if max%2 != 0 {
		max = len(plainText) - 1
	}
	for i := 0; i < max; i += 2 {
		if plainText[i] == plainText[i+1] {
			pos := i + 1
			plainText = plainText[:pos] + "x" + plainText[pos:]
		}
	}

	// pad the plaintext if it has an odd number of letters with a z
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
	new := make([]rune, int(messageLength))

	for i := 0; i < messageLength; i += 2 {
		digraphRectangle := searchForDigraphInKeyTable(kt, msg[i], msg[i+1])

		c1 := digraphRectangle[0]
		c2 := digraphRectangle[1]
		c3 := digraphRectangle[2]
		c4 := digraphRectangle[3]

		// Apply encryption rules
		// If both the letters are in the same row:
		// Take the letter to the right of each one (going back to the leftmost if at the rightmost position).
		if c1 == c3 {
			new[i] = kt[c1][(c2+1)%5]
			new[i+1] = kt[c1][(c4+1)%5]
		} else if c2 == c4 {
			// If both the letters are in the same column:
			// Take the letter below each one (going back to the top if at the bottom).
			new[i] = kt[(c1+1)%5][c2]
			new[i+1] = kt[(c3+1)%5][c2]
		} else { // else, form a rectangle with the two letters and take the letters on the horizontal opposite corner of the rectangle.
			new[i] = kt[c1][c4]
			new[i+1] = kt[c3][c2]
		}
	}

	s := string(new)
	var result string

	for _, r := range s {
		result += string(r)
	}

	return result
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
