package _567

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.True(checkInclusion("ab", "bcdbaooo"))
	//a.False(checkInclusion("ab", "eidboaoo"))
	a.False(checkInclusion("abc", "ccccbbbbaaaa"))
	//a.True(checkInclusion("adc", "dcda"))
	//a.True(checkInclusion("abcdxabcde", "abcdeabcdx"))
}
