package _438

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{0, 101}, findAnagrams(strings.Repeat("a", 200), strings.Repeat("a", 100)))
}
