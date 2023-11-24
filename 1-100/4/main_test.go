package _4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2.0, findMedianSortedArrays([]int{}, []int{2}))
	a.Equal(2.0, findMedianSortedArrays([]int{2}, []int{}))
	a.Equal(1.5, findMedianSortedArrays([]int{1}, []int{2}))
	a.Equal(2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	a.Equal(2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	a.Equal(4.5, findMedianSortedArrays([]int{1, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
