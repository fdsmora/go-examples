package strings

import (
	"strings"
	"unicode"
)

// Leetcode 71. Simplify Path
// https://leetcode.com/problems/simplify-path/
// Medium
type Stack[T any] []T

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	idx := len(*s) - 1
	item := (*s)[idx]
	*s = (*s)[:idx]
	return item, true
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func simplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}
	var currentName strings.Builder
	var dotCount int
	var stack Stack[string]

	for i := 1; i < len(path); i++ {
		ch := rune(path[i])
		switch {
		case unicode.IsLetter(ch) || ch == '_' || unicode.IsDigit(ch):
			if dotCount == 1 {
				dotCount = 0
			}
			currentName.WriteRune(ch)
		case ch == '/':
			handleLastDir(&dotCount, &stack, &currentName)
		case ch == '.':
			dotCount++
			currentName.WriteRune('.')
		}
	}
	handleLastDir(&dotCount, &stack, &currentName)
	currentName.WriteRune('/')
	var notFirst bool
	for _, item := range stack {
		if notFirst {
			currentName.WriteRune('/')
		}
		currentName.WriteString(item)
		notFirst = true
	}
	return currentName.String()
}

func handleLastDir(dotCount *int, stack *Stack[string], sb *strings.Builder) {
	if *dotCount == 1 || *dotCount == 2 {
		if *dotCount == 2 {
			stack.Pop()
		}
		sb.Reset()
		*dotCount = 0
		return
	}
	// if dotCount is 0 or greater than 2
	if sb.Len() > 0 {
		stack.Push(sb.String())
	}
	sb.Reset()
	*dotCount = 0
}

/*

/home/..
       |
S:home,
sb:.
dC:1

////home
        |

	/../
	   |
	S:
	sb:/
	dC:0
		   	   /a/./b/../../c/...././d
		   	                          |
		   	   S:/,c,....,d
		   	   tmp:
		   	   dc:0


		   	   /c/..../d
*/
