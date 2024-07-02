package _115

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, numDistinct("aba", "ab"))
	a.Equal(3, numDistinct("bbb", "bb"))
	a.Equal(3, numDistinct("rabbbit", "rabbit"))
}
