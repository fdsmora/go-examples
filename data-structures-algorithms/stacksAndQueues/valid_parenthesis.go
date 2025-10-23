package stacksandqueues

import (
	"errors"
	"fmt"

	"github.com/fdsmora/go-examples/ds-algorithms/stack"
)

var (
	errStackIsEmpty = errors.New("stack is empty")
)

// LC 20.
func isValidParenthesis(s string) bool {
	if len(s) == 1 {
		return false
	}
	var stack = stack.NewStack()
	chars := make(map[rune]rune, 3)
	chars['{'] = '}'
	chars['['] = ']'
	chars['('] = ')'

	for _, c := range s {
		switch c {
		case '{', '[', '(':
			stack.Push(chars[c])
		case '}', ']', ')':
			cls, err := stack.Pop()
			if errors.Is(err, errStackIsEmpty) {
				fmt.Println("stack is empty")
				return false
			}
			if c != cls {
				return false
			}
		}
	}

	return stack.Size() == 0
}
