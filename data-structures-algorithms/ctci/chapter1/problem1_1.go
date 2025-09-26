package chapter1

// IsUnique returns true if all characters in s are unique.
// This version does not use any extra data structure.
/* func IsUnique(s string) bool {
	if len(s) <= 1 {
		return true
	}

	// Work with runes to handle Unicode
	runes := []rune(s)

	// Sort runes in place
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Check adjacent characters
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}
	return true
} */

// Optimized with bitset, only works when s is only 126 lower-case characters
func IsUniqueLowercase(s string) bool {
	var checker int
	for _, c := range s {
		pos := c - 'a'
		if checker&(1<<pos) > 0 {
			return false
		}
		checker |= 1 << pos
	}
	return true
}
