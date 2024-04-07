package cmd

import "testing"

type testDataVigenere struct {
	key            string
	input          string
	expectedOutput string
}

var testDataMap = []testDataVigenere{
	{
		key:            "run",
		input:          "test string!",
		expectedOutput: "kyfk mgicax!",
	},
	{
		key:            "lorem ipsum",
		input:          "scelerisque fringilla turpis a euismod",
		expectedOutput: "dqvpqzxkkgp timzoxdfm eiitua p woudafh",
	},
	{
		key:            "maximus",
		input:          "molestie nunc placerat",
		expectedOutput: "yoimenaq nrvo jdmcbzmn",
	},
	{
		key:            "scelerisque",
		input:          "nec luctus",
		expectedOutput: "fgg wytbmi",
	},
	{
		key:            "semper",
		input:          "consectetur sodales lectus sem cursus",
		expectedOutput: "uszhitlifjv jghmaij dioiyj kiy ryikye",
	},
	{
		key:            "pellentesque",
		input:          "quis sem scelerisque",
		expectedOutput: "fytd wrf wuufigmdbyr",
	},
	{
		key:            "lobortis diam",
		input:          "enim vitae tempus venenatis",
		expectedOutput: "pbja mbbsh beyait jvgmfdbie",
	},
}

func TestShiftTextByKeywordEncode(t *testing.T) {
	for _, val := range testDataMap {
		msg := ShiftTextByKeyword(val.input, false, val.key)
		if msg != val.expectedOutput {
			t.Fatalf(`encode ShiftText("%v") returned %q, should have been "%v"`, val.input, msg, val.expectedOutput)
		}
	}
}

func TestShiftTextByKeywordDecode(t *testing.T) {
	for _, val := range testDataMap {
		msg := ShiftTextByKeyword(val.expectedOutput, true, val.key)
		if msg != val.input {
			t.Fatalf(`decode ShiftText("%v") returned %q, should have been "%v"`, val.expectedOutput, msg, val.input)
		}
	}
}
