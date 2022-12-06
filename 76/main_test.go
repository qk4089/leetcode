package _76

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("bc", minWindow("abc", "bc"))
	a.Equal("BANC", minWindow("ADOBECODEBANC", "ABC"))
	a.Equal("a", minWindow("a", "a"))
	a.Equal("", minWindow("a", "aa"))
}
