package stacksandqueues

import (
	"github.com/fdsmora/go-examples/ds-algorithms/stack"
)

// LC 20.
// my solution, O(n) & O(n) space
func isValidParenthesis(s string) bool {
	if len(s) == 1 {
		return false
	}
	var stack = stack.Stack[rune]{}
	chars := map[rune]rune{'{': '}', '[': ']', '(': ')'}

	for _, c := range s {
		switch c {
		case '{', '[', '(':
			stack.Push(chars[c])
		case '}', ']', ')':
			if cls, ok := stack.Pop(); !ok || cls != c {
				return false
			}
		}
	}

	return stack.Size() == 0
}

// Solution from neetcode
/*func isValidParenthesis(s string) bool {
	if len(s) == 1 {
		return false
	}
	var stack = stack.Stack[rune]{}
	closeToOpen := map[rune]rune{'}': '{', ']': '[', ')': '('}

	for _, c := range s {
		if open, exists := closeToOpen[c]; exists {
			if !stack.IsEmpty() {
				top, ok := stack.Pop()
				if !ok || top != open {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(c)
		}
	}
	return stack.Size() == 0
}*/
