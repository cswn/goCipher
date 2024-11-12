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
		expectedOutput: "anrlfmudrmpsrhxblzmv",
	},
	{
		key:            "scelerisque",
		input:          "nec luctus",
		expectedOutput: "ocerqeobev",
	},
	{
		key:            "semper",
		input:          "consectetur sodales lectus sem cursus",
		expectedOutput: "bqvambopunsetbfgmehrdqnremcizfrnmv",
	},
	{
		key:            "pellentesque",
		input:          "quis sem scelerisque",
		expectedOutput: "uavcqpiudpnlikquql",
	},
	{
		key:            "lobortis diam",
		input:          "enim vitae tempus venenatis",
		expectedOutput: "nwsilcrmhohsqkiwnwnwmrsd",
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
