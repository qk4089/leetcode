package _30

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{0, 1}, findSubstring("baabaab", []string{"ab", "ba", "aa"}))
	a.Equal([]int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	a.Equal([]int{}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	a.Equal([]int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
