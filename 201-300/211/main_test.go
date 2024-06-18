package _211

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_d(t *testing.T) {
	a := assert.New(t)
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	a.False(wordDictionary.Search("pad"))
	a.True(wordDictionary.Search("bad"))
	a.True(wordDictionary.Search(".ad"))
	a.True(wordDictionary.Search("b.."))
}
