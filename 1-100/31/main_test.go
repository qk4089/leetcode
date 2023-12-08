package _31

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)

	arr3 := []int{1, 1}
	nextPermutation(arr3)
	a.Equal([]int{1, 1}, arr3)

	arr2 := []int{1, 1, 5}
	nextPermutation(arr2)
	a.Equal([]int{1, 5, 1}, arr2)
	nextPermutation(arr2)
	a.Equal([]int{5, 1, 1}, arr2)
	nextPermutation(arr2)
	a.Equal([]int{1, 1, 5}, arr2)

	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 4, 5, 7, 6}, arr1)
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 4, 6, 5, 7}, arr1)
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 4, 6, 7, 5}, arr1)
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 4, 7, 5, 6}, arr1)
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 4, 7, 6, 5}, arr1)
	nextPermutation(arr1)
	a.Equal([]int{1, 2, 3, 5, 4, 6, 7}, arr1)
}
