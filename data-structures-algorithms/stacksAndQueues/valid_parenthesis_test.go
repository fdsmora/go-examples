package stacksandqueues

import "testing"

func TestIsValidParenthesis(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"simple pair", "()", true},
		{"closing pair", "]]", false},
		{"multiple pairs", "()[]{}", true},
		{"nested", "({[]})", true},
		{"complex nested", "(([]){})", true},
		{"unmatched closing", ")", false},
		{"unmatched opening", "(", false},
		{"wrong order", "(]", false},
		{"crossed", "([)]", false},
		{"missing closing", "([", false},
		{"extra closing in middle", "{[]} )", false},
		{"long valid", "(((((((((())))))))))", true},
		{"long invalid", "(((((((()))))))) )", false},
	}

	for _, tt := range tests {
		if got := isValidParenthesis(tt.input); got != tt.want {
			t.Errorf("%s: isValidParenthesis(%q) = %v; want %v", tt.name, tt.input, got, tt.want)
		}
	}
}
