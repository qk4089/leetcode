package _84

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(3, largestRectangleArea([]int{2, 1, 2}))
	a.Equal(6, largestRectangleArea([]int{4, 2, 0, 3, 2, 5}))
	a.Equal(10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
