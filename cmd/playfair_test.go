package cmd

import "testing"

type testDataPlayfair struct {
	key                string
	input              string
	expectedOutput     string
	expectedDecryption string
}

var testDataPlayfairMap = []testDataPlayfair{
	{
		key:                "run",
		input:              "test string!",
		expectedOutput:     "qgtotouhbe",
		expectedDecryption: "teststring",
	},
	{
		key:                "lorem ipsum",
		input:              "scelerisque fringilla turpis euismod",
		expectedOutput:     "pdmomepuyfuqlstdbimiqaospuufpulrnr",
		expectedDecryption: "scelerisquefringillaturpiseuismodx",
	},
	{
		key:                "maximus",
		input:              "molestie nunc placerat",
		expectedOutput:     "anrlfmudrmpsrhxblzmv",
		expectedDecryption: "molestienuncplacerat",
	},
	{
		key:                "scelerisque",
		input:              "nec luctus",
		expectedOutput:     "ocerqeobev",
		expectedDecryption: "necluctusx",
	},
	{
		key:                "semper",
		input:              "consectetur sodales lectus sem cursus",
		expectedOutput:     "bqvambopunsetbfgmehrdqnremcizfrnmv",
		expectedDecryption: "consectetursodaleslectussemcursusx",
	},
	{
		key:                "pellentesque",
		input:              "quis sem scelerisque",
		expectedOutput:     "uavcqpiudpnlikquql",
		expectedDecryption: "quissemscelerisque",
	},
	{
		key:                "lobortis diam",
		input:              "enim vitae tempus venenatis",
		expectedOutput:     "nwsilcrmhohsqkiwnwnwmrsd",
		expectedDecryption: "enimvitaetempusvenenatis",
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

func TestShiftTextByDigraphDecode(t *testing.T) {
	for _, val := range testDataPlayfairMap {
		msg := ShiftTextByDigraph(val.expectedOutput, true, val.key)
		if msg != val.expectedDecryption {
			t.Fatalf(`decode ShiftTextByDigraph("%v") returned %q, should have been "%v"`, val.expectedOutput, msg, val.expectedDecryption)
		}
	}
}
