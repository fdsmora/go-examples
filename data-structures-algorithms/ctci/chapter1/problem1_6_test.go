package chapter1

import "testing"

func TestCompress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abcdef", "abcdef"}, // no compression
		{"aabbcc", "a2b2c2"}, // no compression
		{"aaa", "a3"},
		{"a", "a"},
		{"", ""},
		{"aaAA", "a2A2"},
		{"abcdee", "abcdee"},
	}

	for _, tt := range tests {
		result := CompressBetter(tt.input)
		if result != tt.expected {
			t.Errorf("Compress(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
