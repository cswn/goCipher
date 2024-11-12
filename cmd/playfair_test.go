package cmd

import "testing"

type testDataPlayfair struct {
	key            string
	input          string
	expectedOutput string
}

var testDataPlayfairMap = []testDataPlayfair{
	{
		key:            "run",
		input:          "test string!",
		expectedOutput: "qgtotouhbe",
	},
	{
		key:            "lorem ipsum",
		input:          "scelerisque fringilla turpis a euismod",
		expectedOutput: "pdmomepuyfuqlstdbimiqaospuumaparrc",
	},
	{
		key:            "maximus",
		input:          "molestie nunc placerat",
		expectedOutput: "nppatukdpsshlmbdbudq ",
	},
	{
		key:            "scelerisque",
		input:          "nec luctus",
		expectedOutput: "pcanseuqtr",
	},
	{
		key:            "semper",
		input:          "consectetur sodales lectus sem cursus",
		expectedOutput: "dnsxaduduqsttifqcupadsqtucnbqstqtr",
	},
	{
		key:            "pellentesque",
		input:          "quis sem scelerisque",
		expectedOutput: "rqhtucnrdapatgtrzk",
	},
	{
		key:            "lobortis diam",
		input:          "enim vitae tempus venenatis",
		expectedOutput: "cpgoyfqddubpuzqxcpcpdqht",
	},
}

func TestShiftTextByDigraphEncrypt(t *testing.T) {
	for _, val := range testDataPlayfairMap {
		msg := ShiftTextByDigraph(val.input, false, val.key)
		if msg != val.expectedOutput {
			t.Fatalf(`encode ShiftTextByDigraph("%v") returned %q, should have been "%v"`, val.input, msg, val.expectedOutput)
		}
	}
}
