package miscelaneous

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name,
		input,
		want string
	}{
		{
			"repeated chars",
			"ccc",
			"ccc",
		},
		{
			"longer string",
			"hellosannasmith",
			"sannas",
		},
		{
			"longer repeated",
			"aaaaaaaaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		{
			"longer repeated but shorter",
			"aaaaaaaaaaaaaaaabaaaaaaa",
			"aaaaaaaaaaaaaaaa",
		},
		{
			"2 chars",
			"xx",
			"xx",
		},
		{
			"1 char",
			"x",
			"x",
		},
		{
			"another string",
			"abracecars",
			"racecar",
		},
		{
			"even length",
			"abbaabcd",
			"abba",
		},
		{
			"bananas",
			"bananas",
			"anana",
		},
		{
			"long symmetric",
			"forgeeksskeegfor",
			"geeksskeeg",
		},
		{
			"empty string",
			"",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("test '%s' failed. Got:'%v', want:'%v'", tt.name, got, tt.want)
			}
		})
	}
}
