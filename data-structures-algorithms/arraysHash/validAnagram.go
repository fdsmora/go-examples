package arrayshash

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := [26]int{}

	for i, ch := range s {
		charCount[ch-'a']++
		charCount[t[i]-'a']--
	}

	for _, c := range charCount {
		if c != 0 {
			return false
		}
	}
	return true
}
