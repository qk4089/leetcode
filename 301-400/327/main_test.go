package _327

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, countRangeSum([]int{1, 1, 2}, 0, 2))
	//a.Equal(3, countRangeSum([]int{-2, 5, -1}, -2, 2))
	//a.Equal(1, countRangeSum([]int{0}, 0, 0))
}
