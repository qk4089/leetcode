package _9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(isPalindrome(0))
	a.True(isPalindrome(8))
	a.False(isPalindrome(10))
	a.True(isPalindrome(121))
	a.True(isPalindrome(1221))
}
