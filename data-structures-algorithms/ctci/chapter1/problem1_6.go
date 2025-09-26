package chapter1

import (
	"strconv"
	"strings"
)

func Compress(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var sb strings.Builder
	sb.Grow(2 * n) // worst case: every char unique → output length = 2n, e.g. "abc" -> "a1b1c1"
	i, count := 1, 1
	for ; i < n; i++ {
		if s[i] != s[i-1] {
			sb.WriteByte(s[i-1]) // Assumes s has only ASCII characters, if it had Unicode, use runes (sb.WriteRune)
			sb.WriteString(strconv.Itoa(count))
			count = 1
		} else {
			count++
		}
	}
	sb.WriteByte(s[i-1])
	sb.WriteString(strconv.Itoa(count))

	if sb.Len() <= n {
		return sb.String()
	}
	return s
}
