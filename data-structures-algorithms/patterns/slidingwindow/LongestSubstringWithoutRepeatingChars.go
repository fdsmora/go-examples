package slidingwindow

import (
	"math"
)

// leetcode problem 3
// Runtime complexity: O(n+n) -> O(n), as e iterates all characters and when duplicate if found, oldS
// traverses all characters until s, deleting each one of them from the hash map
// space complexity: O(n) as the hash map can hold all of the characters if all are unique
func LengthOfLongestSubstring(str string) int {
	if len(str) <= 1 {
		return len(str)
	}
	var max, s, oldS, e int = 1, 0, 0, 0
	var ch rune
	hm := make(map[rune]int)
	for e, ch = range str {
		if _, ok := hm[ch]; !ok {
			hm[ch] = e // store the index of the character
		} else {
			max = int(math.Max(float64(max), float64(e-s)))
			duplicateIdx := hm[ch]
			s = duplicateIdx + 1
			for oldS != s {
				delete(hm, rune(str[oldS]))
				oldS++
			}
			hm[ch] = e
		}
	}

	return int(math.Max(float64(max), float64(e-s+1)))
}
