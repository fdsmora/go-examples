package chapter1

// modify s in place, so pass s as []rune
func Replace(s []rune, m int) []rune {
	k := len(s) - 1 // last position in s
	for i := m - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s[k] = '0'
			k--
			s[k] = '2'
			k--
			s[k] = '%'
		} else {
			s[k] = s[i]
		}
		k--
	}
	return s
}
