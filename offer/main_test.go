package offer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_51(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, reversePairs([]int{7, 5, 6, 4}))
	a.Equal(0, reversePairs([]int{4, 5, 6, 7}))
	a.Equal(4, reversePairs([]int{1, 3, 2, 3, 1}))
	a.Equal(10, reversePairs([]int{5, 4, 3, 2, 1}))
	a.Equal(6, reversePairs([]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}))
}
