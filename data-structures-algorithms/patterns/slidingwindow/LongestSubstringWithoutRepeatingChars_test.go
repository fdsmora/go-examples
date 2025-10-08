package slidingwindow

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},    // "b"
		{"pwwkew", 3},   // "wke"
		{"", 0},         // empty string
		{"a", 1},        // single char
		{"au", 2},       // "au"
		{"dvdf", 3},     // "vdf"
		{"anviaj", 5},   // "nviaj"
		{"tmmzuxt", 5},  // "mzuxt"
		{"abba", 2},     // "ab" or "ba"
		{"abcdefg", 7},  // all unique
		{"aab", 2},      // "ab"
		{"abcadef", 6},  // "bcadef",
	}

	for _, tt := range tests {
		result := LengthOfLongestSubstring(tt.input)
		if result != tt.expected {
			t.Errorf("LengthOfLongestSubstring(%q) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
