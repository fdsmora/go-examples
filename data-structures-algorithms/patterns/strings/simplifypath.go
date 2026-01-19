package strings

import (
	"strings"
	"unicode"
)

// Leetcode 71. Simplify Path
// https://leetcode.com/problems/simplify-path/
// Medium
func simplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}
	var final, tmp strings.Builder
	var dotCount int
	var inWord bool

	for _, ch := range path {
		if ch == '/' {
			// skip the '/'
			inWord = false
		} else if ch == '_' || unicode.IsLetter(ch) {
			if inWord {
				tmp.WriteRune(ch)
				continue
			} else {
				// path[i] is in the first character after a '/'
				switch {
				case dotCount <= 1:
					final.WriteString(tmp.String()) // flush tmp
					tmp.Reset()
					tmp.WriteString("/" + string(ch))
					dotCount = 0
				case dotCount == 2:
					tmp.Reset() // discard the last directory
					tmp.WriteString("/" + string(ch))
				case dotCount > 2:
					final.WriteString("/" + strings.Repeat(".", dotCount))
					tmp.Reset()
					tmp.WriteString("/" + string(ch))
				}
				inWord = true
				dotCount = 0
			}
		} else if ch == '.' {
			dotCount++
		}
	}
	final.WriteString(tmp.String())
	if final.Len() == 0 {
		return "/"
	}
	return final.String()
}
