package _57

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{5, 7}}, insert([][]int{}, []int{5, 7}))
	a.Equal([][]int{{0, 5}}, insert([][]int{{1, 5}}, []int{0, 3}))
	a.Equal([][]int{{1, 2}, {5, 6}, {9, 10}}, insert([][]int{{1, 2}, {9, 10}}, []int{5, 6}))
	a.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
