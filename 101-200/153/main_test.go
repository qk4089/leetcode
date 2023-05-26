package _153

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, findMin([]int{3, 4, 5, 1, 2}))
	a.Equal(0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	a.Equal(11, findMin([]int{11, 13, 15, 17}))
	a.Equal(1, findMin([]int{2, 1}))
	a.Equal(1, findMin([]int{1, 2}))
	a.Equal(1, findMin([]int{5, 1, 2, 3, 4}))
}
