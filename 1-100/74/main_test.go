package _74

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	a.False(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	a.True(searchMatrix([][]int{{-8, -7, -5, -3, -3, -1, 1}, {2, 2, 2, 3, 3, 5, 7}, {8, 9, 11, 11, 13, 15, 17}, {18, 18, 18, 20, 20, 20, 21}, {23, 24, 26, 26, 26, 27, 27}, {28, 29, 29, 30, 32, 32, 34}}, -5))
}
