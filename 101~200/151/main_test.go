package _151

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("blue is sky the", reverseWords("the sky is blue"))
	a.Equal("world hello", reverseWords("  hello world  "))
}
