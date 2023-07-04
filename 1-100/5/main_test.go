package _5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("bdb", longestPalindrome("cbdba"))
	//fmt.Print(getPalindrome("aba", 1, 1))
}
