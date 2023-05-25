package _15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal([][]int{{3, 5}}, two([]int{3, 3, 4, 4, 5, 5}, 8))
	//a.Equal([][]int{{4, 4, -8}, {3, 5, -8}}, threeSum([]int{-8, 3, 3, 4, 4, 5, 5}))
	a.Equal([][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
	//a.Equal([][]int{{0, 0, 0}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//a.Equal([][]int{{0, 0, 0}}, threeSum([]int{3, 0, -2, -1, 1, 2}))
}
