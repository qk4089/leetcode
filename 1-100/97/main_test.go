package _97

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.False(isInterleave("a", "b", "a"))
	a.True(isInterleave("ab", "bc", "babc"))
	a.True(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	a.False(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	a.False(isInterleave("ababbb", "baaaba", "babababababb"))
	a.False(isInterleave("abababababababababababababababababababababababababababababababababababababababababababababababababbb", "babababababababababababababababababababababababababababababababababababababababababababababababaaaba", "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"))
}
