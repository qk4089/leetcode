package _493

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_51(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, reversePairs([]int{-5, -5}))
	a.Equal(2, reversePairs([]int{1, 3, 2, 3, 1}))
	a.Equal(3, reversePairs([]int{2, 4, 3, 5, 1}))
	a.Equal(9, reversePairs([]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}))
	a.Equal(9, reversePairs([]int{-185, 143, -154, -338, -269, 287}))
}
