package _3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(1, lengthOfLongestSubstring(" "))
	//a.Equal(2, lengthOfLongestSubstring("aab"))
	a.Equal(3, lengthOfLongestSubstring("abcabcbb"))
}
