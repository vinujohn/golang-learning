package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isBalanced(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(', '{', '[', '<':
			stack = append(stack, c)
		case ')', '}', ']', '>':
			if len(stack) == 0 {
				return false
			}
			if (c == ')' && stack[len(stack)-1] != '(') ||
				(c == '}' && stack[len(stack)-1] != '{') ||
				(c == ']' && stack[len(stack)-1] != '[') ||
				(c == '>' && stack[len(stack)-1] != '<') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestBalancedParens(t *testing.T) {
	assert.True(t, isBalanced("(hello)world"))
	assert.False(t, isBalanced("hello)world("))
	assert.True(t, isBalanced("hello{world}abc"))
	assert.False(t, isBalanced("hello{world}abc}def"))
	assert.False(t, isBalanced("world>"))
	assert.True(t, isBalanced("helloabc"))
}
