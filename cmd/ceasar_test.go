package cmd

import "testing"

//var testDataMap map[int][]string

// hashmap including input, key, expected output
// then iterate over the map using the values in each test case

type testData struct {
	key            int64
	input          string
	expectedOutput string
}

var testDataMap = []testData{
	{
		key:            2,
		input:          "test string",
		expectedOutput: "vguv uvtkpi",
	},
	{
		key:            4,
		input:          "lorem ipsum",
		expectedOutput: "psviq mtwyq",
	},
	{
		key:            8,
		input:          "dolor sit amet",
		expectedOutput: "lwtwz aqb iumb",
	},
	{
		key:            12,
		input:          "sed sapien et",
		expectedOutput: "eqp embuqz qf",
	},
	{
		key:            13,
		input:          "feugiat accumsan",
		expectedOutput: "srhtvng npphzfna",
	},
	{
		key:            17,
		input:          "pellentesque",
		expectedOutput: "gvccvekvjhlv",
	},
	{
		key:            23,
		input:          "lobortis diam",
		expectedOutput: "ilyloqfp afxj",
	},
}

func TestShiftTextEncode(t *testing.T) {
	for _, val := range testDataMap {
		msg := ShiftText(val.input, false, val.key)
		if msg != val.expectedOutput {
			t.Fatalf(`encode ShiftText("%v") returned %q, should have been "%v"`, val.input, msg, val.expectedOutput)
		}
	}
}

func TestShiftTextDecode(t *testing.T) {
	for _, val := range testDataMap {
		msg := ShiftText(val.expectedOutput, true, val.key)
		if msg != val.input {
			t.Fatalf(`decode ShiftText("%v") returned %q, should have been "%v"`, val.expectedOutput, msg, val.input)
		}
	}
}
