package chapter1

import "sort"

func IsPermutationHT(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	hm := make(map[rune]int)
	for _, c := range a {
		hm[c]++
	}
	for _, c := range b {
		hm[c]--
		if hm[c] < 0 {
			return false
		}
	}
	return true
}

// Asuming only ASCII characters
func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	occur := make([]int, 128) // use int, not byte, to avoid overflow

	for i := 0; i < len(a); i++ {
		occur[a[i]]++
	}

	for i := 0; i < len(b); i++ {
		occur[b[i]]--
		if occur[b[i]] < 0 {
			return false
		}
	}

	return true
}

// runtime O(nlogn), O(n) space
func IsPermutationONLogN(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aR := []rune(a)
	bR := []rune(b)

	sort.Slice(aR, func(i, j int) bool {
		return aR[i] < aR[j]
	})
	sort.Slice(bR, func(i, j int) bool {
		return bR[i] < bR[j]
	})

	return string(aR) == string(bR)
}
