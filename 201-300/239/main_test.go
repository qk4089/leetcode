package _239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	a.Equal([]int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
	a.Equal([]int{10, 10, 9, 2}, maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))
}
