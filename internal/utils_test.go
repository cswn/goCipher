package internal

import (
	"testing"
	"unicode"
)

type testDataFilterString struct {
	input          string
	callback       func(rune) bool
	expectedOutput string
}

type testDataRuneCount struct {
	input          string
	expectedOutput int
}

var testDataFilterStringMap = []testDataFilterString{
	{
		input: "aaaaaBBccc",
		callback: func(letter rune) bool {
			return unicode.IsLower(letter)
		},
		expectedOutput: "aaaaaccc",
	},
	{
		input: "lorem ipsum sit dolorem em",
		callback: func(letter rune) bool {
			return letter != 'm'
		},
		expectedOutput: "lore ipsu sit dolore e",
	},
	{
		input: "$aR3411ys3cur3p455w0rd!",
		callback: func(letter rune) bool {
			return unicode.IsLetter(letter)
		},
		expectedOutput: "aRyscurpwrd",
	},
}

var testDataRuneCountMap = []testDataRuneCount{
	{
		input:          "lorem ipsum",
		expectedOutput: 10,
	},
	{
		input:          " ",
		expectedOutput: 0,
	},
	{
		input:          "...",
		expectedOutput: 0,
	},
	{
		input:          "AbCdEfGhIjKlMnOpQrStUvWxYz",
		expectedOutput: 26,
	},
}

func TestFilterString(t *testing.T) {
	for _, val := range testDataFilterStringMap {
		msg := FilterString(val.input, val.callback)
		if msg != val.expectedOutput {
			t.Fatalf(`FilterString("%v") returned %q, should have been "%v"`, val.input, msg, val.expectedOutput)
		}
	}
}

func TestRuneCountInStringLettersOnly(t *testing.T) {
	for _, val := range testDataRuneCountMap {
		msg := RuneCountInStringLettersOnly(val.input)
		if msg != val.expectedOutput {
			t.Fatalf(`RuneCountInStringLettersOnly("%v") returned %q, should have been "%v"`, val.input, msg, val.expectedOutput)
		}
	}
}
