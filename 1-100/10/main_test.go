package _10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.False(isMatch("aa", "a"))
	a.True(isMatch("aa", "a*"))
	a.True(isMatch("ab", ".*"))
	a.True(isMatch("a", "a*a"))
	a.False(isMatch("aaa", "aaaa"))
	a.True(isMatch("a", "ab*"))
	a.True(isMatch("aab", "a*b"))
	a.True(isMatch("aab", "c*a*b"))
	a.False(isMatch("ab", ".*c"))
	a.True(isMatch("aaa", "a*a"))
	a.False(isMatch("mississippi", "mis*is*p*."))
	a.False(isMatch("aaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*"))
}
