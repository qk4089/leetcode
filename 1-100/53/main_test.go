package _53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	a.Equal(23, maxSubArray([]int{5, 4, -1, 7, 8}))
	a.Equal(3, maxSubArray([]int{1, 2}))
}
