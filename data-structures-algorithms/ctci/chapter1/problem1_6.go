package chapter1

import (
	"strconv"
	"strings"
)

// My solution
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

// CTCI solution, better because it calculates the final length first
func CompressBetter(s string) string {
	finalLength := countCompression(s)
	if finalLength > len(s) {
		return s
	}

	var sb strings.Builder
	sb.Grow(finalLength)
	consecutiveCount := 0
	for i := 0; i < len(s); i++ {
		consecutiveCount++
		if i+1 >= len(s) || s[i] != s[i+1] {
			sb.WriteByte(s[i])
			sb.WriteString(strconv.Itoa(consecutiveCount))
			consecutiveCount = 0
		}
	}
	return sb.String()
}

func countCompression(s string) int {
	var compressLength, countConsecutive int
	for i := 0; i < len(s); i++ {
		countConsecutive++
		if i+1 >= len(s) || s[i] != s[i+1] {
			compressLength += 1 + len(strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}
	}
	return compressLength
}
