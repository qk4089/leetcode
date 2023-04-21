package _154

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, findMin([]int{1, 3}))
	a.Equal(1, findMin([]int{3, 1}))
	a.Equal(1, findMin([]int{1, 3, 1}))
	a.Equal(1, findMin([]int{1, 3, 5}))
	a.Equal(0, findMin([]int{2, 2, 2, 0, 1}))
	a.Equal(1, findMin([]int{1, 3, 3, 3}))
	a.Equal(1, findMin([]int{3, 3, 1, 3}))
}
