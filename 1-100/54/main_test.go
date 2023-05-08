package _54

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{1, 2, 3}, spiralOrder([][]int{{1}, {2}, {3}}))
	a.Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	a.Equal([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	a.Equal([]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}, spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}))
}
