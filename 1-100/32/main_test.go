package _32

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, longestValidParentheses("(()"))
	a.Equal(4, longestValidParentheses("()()"))
	a.Equal(4, longestValidParentheses(")()())"))
	a.Equal(6, longestValidParentheses(")())(()())"))
	a.Equal(2, longestValidParentheses("()(()"))
}
