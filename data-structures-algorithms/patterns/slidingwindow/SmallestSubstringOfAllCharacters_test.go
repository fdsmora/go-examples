package slidingwindow

import (
	"testing"
)

func TestSmallestSubstringAllChars(t *testing.T) {
	tests := map[string]struct {
		inputChars []string
		inputStr,
		want string
	}{
		"ok": {
			[]string{"a", "b", "c", "d"},
			"xadcbamiabrvddd",
			"ab",
		},
		"no match": {
			[]string{"a", "b", "c", "d"},
			"xyzmukliew",
			"",
		},
		"s": {
			[]string{"a", "b", "c", "d"},
			"xyzmuklaiew",
			"a",
		},
		"one match": {
			[]string{"a"},
			"a",
			"a",
		},
		"empty string": {
			[]string{"a", "b", "c"},
			"",
			"",
		},
		"empty chars": {
			[]string{},
			"abcdef",
			"",
		},
		"single char in string": {
			[]string{"a", "b", "c"},
			"x",
			"",
		},
		"all chars match": {
			[]string{"a", "b", "c"},
			"abc",
			"abc",
		},
		"consecutive matches at start": {
			[]string{"a", "b"},
			"abxyz",
			"ab",
		},
		"consecutive matches at end": {
			[]string{"a", "b"},
			"xyzab",
			"ab",
		},
		"consecutive matches in middle": {
			[]string{"a", "b"},
			"xyzabpqr",
			"ab",
		},
		"single char repeated": {
			[]string{"a"},
			"xaaay",
			"aaa",
		},
		"multiple valid substrings, return shortest": {
			[]string{"a", "b", "c"},
			"xabcyabc",
			"abc",
		},
		"longer valid substring at start": {
			[]string{"a", "b"},
			"aabbxab",
			"ab",
		},
		"all string matches": {
			[]string{"a", "b", "c", "d", "e"},
			"abcde",
			"abcde",
		},
		"alternating valid/invalid": {
			[]string{"a", "b"},
			"axbxaxb",
			"a",
		},
		"valid substring at boundaries": {
			[]string{"a"},
			"xax",
			"a",
		},
		"multiple single chars": {
			[]string{"a", "b", "c"},
			"xaybzc",
			"a",
		},
		"long valid substring vs short one": {
			[]string{"a", "b", "c"},
			"abcabcabcxabxabcabc",
			"ab",
		},
		"chars with repetition in pattern": {
			[]string{"a", "a", "b"}, // duplicate 'a' in chars
			"xabay",
			"aba",
		},
		"only invalid chars": {
			[]string{"a", "b"},
			"xyz",
			"",
		},
		"one valid char at start": {
			[]string{"a", "b", "c"},
			"axyz",
			"a",
		},
		"one valid char at end": {
			[]string{"a", "b", "c"},
			"xyza",
			"a",
		},
		"mixed case sensitivity": {
			[]string{"a", "b"},
			"ABab",
			"ab",
		},
		"special characters": {
			[]string{"!", "@", "#"},
			"x!@#y",
			"!@#",
		},
		"numbers as strings": {
			[]string{"1", "2", "3"},
			"x123y1",
			"1",
		},
		"long string with short result": {
			[]string{"a"},
			"xxxxxxxxxxxxxxxxxxxxxxxxa",
			"a",
		},
		"entire string is valid": {
			[]string{"a"},
			"aaaaaaaaaa",
			"aaaaaaaaaa",
		},
		"pattern at multiple positions same length": {
			[]string{"a", "b"},
			"xabyab",
			"ab",
		},
		"substring with all different chars": {
			[]string{"a", "b", "c", "d", "e"},
			"xabcdey",
			"abcde",
		},
		"two char pattern": {
			[]string{"x", "y"},
			"axyb",
			"xy",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SmallestSubstringAllChars(tc.inputChars, tc.inputStr)
			if got != tc.want {
				t.Errorf("Got: '%s', want: '%s'", got, tc.want)
			}
		})
	}
}
