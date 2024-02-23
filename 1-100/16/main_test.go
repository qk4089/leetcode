package _16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(-101, threeSumClosest([]int{-100, -98, -2, -1}, -101))
	a.Equal(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	a.Equal(0, threeSumClosest([]int{0, 0, 0}, 1))
}
