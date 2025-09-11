package cmd

import "testing"

type testDataSBX struct {
	key            uint
	input          string
	expectedOutput string
}

var testDataSBXMap = []testDataSBX{
	{
		key:            54,
		input:          "9f8a8a9f9d959f8a9a9f8990",
		expectedOutput: "a9bcbca9aba3a9bcaca9bfa6",
	},
	{
		key:            9,
		input:          "646566656e64746865726964676561746475736b",
		expectedOutput: "6d6c6f6c676d7d616c7b606d6e6c687d6d7c7a62",
	},
	{
		key:            222,
		input:          "4c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e73656374657475722061646970697363696e6720656c69742c2073656420646f20656975736d6f642074656d706f7220696e6369646964756e74207574206c61626f726520657420646f6c6f7265206d61676e6120616c697175612e20557420656e696d206164206d696e696d2076656e69616d2c2071756973206e6f737472756420657865726369746174696f6e20756c6c616d636f206c61626f726973206e69736920757420616c697175697020657820656120636f6d6d6f646f20636f6e7365717561742e2044756973206175746520697275726520646f6c6f7220696e20726570726568656e646572697420696e20766f6c7570746174652076656c697420657373652063696c6c756d20646f6c6f726520657520667567696174206e756c6c612070617269617475722e204578636570746575722073696e74206f6363616563617420637570696461746174206e6f6e2070726f6964656e742c2073756e7420696e2063756c706120717569206f666669636961206465736572756e74206d6f6c6c697420616e696d20696420657374206c61626f72756d",
		expectedOutput: "92b1acbbb3feb7aeadabb3febab1b2b1acfeadb7aafebfb3bbaaf2febdb1b0adbbbdaabbaaabacfebfbab7aeb7adbdb7b0b9febbb2b7aaf2feadbbbafebab1febbb7abadb3b1bafeaabbb3aeb1acfeb7b0bdb7bab7baabb0aafeabaafeb2bfbcb1acbbfebbaafebab1b2b1acbbfeb3bfb9b0bffebfb2b7afabbff0fe8baafebbb0b7b3febfbafeb3b7b0b7b3fea8bbb0b7bfb3f2feafabb7adfeb0b1adaaacabbafebba6bbacbdb7aabfaab7b1b0feabb2b2bfb3bdb1feb2bfbcb1acb7adfeb0b7adb7feabaafebfb2b7afabb7aefebba6febbbffebdb1b3b3b1bab1febdb1b0adbbafabbfaaf0fe9aabb7adfebfabaabbfeb7acabacbbfebab1b2b1acfeb7b0feacbbaeacbbb6bbb0babbacb7aafeb7b0fea8b1b2abaeaabfaabbfea8bbb2b7aafebbadadbbfebdb7b2b2abb3febab1b2b1acbbfebbabfeb8abb9b7bfaafeb0abb2b2bffeaebfacb7bfaaabacf0fe9ba6bdbbaeaabbabacfeadb7b0aafeb1bdbdbfbbbdbfaafebdabaeb7babfaabfaafeb0b1b0feaeacb1b7babbb0aaf2feadabb0aafeb7b0febdabb2aebffeafabb7feb1b8b8b7bdb7bffebabbadbbacabb0aafeb3b1b2b2b7aafebfb0b7b3feb7bafebbadaafeb2bfbcb1acabb3",
	},
}

func TestSingleByteXOREncode(t *testing.T) {
	for _, val := range testDataSBXMap {
		hexResult := SingleByteXor(val.input, val.key)
		if hexResult != val.expectedOutput {
			t.Fatalf(`encode SingleByteXOR("%v") returned %q, should have been "%v"`, val.input, hexResult, val.expectedOutput)
		}
	}
}

func TestSingleByteXORDecode(t *testing.T) {
	for _, val := range testDataSBXMap {
		hexResult := SingleByteXor(val.expectedOutput, val.key)
		if hexResult != val.input {
			t.Fatalf(`decode SingleByteXOR("%v") returned %q, should have been "%v"`, val.expectedOutput, hexResult, val.input)
		}
	}
}
