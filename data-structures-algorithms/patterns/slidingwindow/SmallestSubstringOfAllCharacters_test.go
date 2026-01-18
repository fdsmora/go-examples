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
