package _316

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", removeDuplicateLetters("bcabc"))
	a.Equal("acdb", removeDuplicateLetters("cbacdcbc"))
	a.Equal("bac", removeDuplicateLetters("bbcaac"))
	a.Equal("abc", removeDuplicateLetters("abacb"))
}
