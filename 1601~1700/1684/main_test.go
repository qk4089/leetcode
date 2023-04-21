package _1684

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(2, main_2("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	//a.Equal(7, main_2("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	a.Equal(4, main_2("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
