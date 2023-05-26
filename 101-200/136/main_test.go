package _136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(-1, singleNumber([]int{2, -1, 2}))
	a.Equal(4, singleNumber([]int{4, 1, 2, 1, 2}))
	a.Equal(1, singleNumber([]int{1}))
	a.Equal(-2, singleNumber([]int{-1, -1, -2, 1, 1}))
}
