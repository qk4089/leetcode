package _60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, subarraySum([]int{1, 1, 1}, 2))
	a.Equal(2, subarraySum([]int{1, 1, 2}, 2))
	a.Equal(3, subarraySum([]int{1, 2, 1, 2}, 3))
	a.Equal(2, subarraySum([]int{1, 2, 1, 2}, 2))
	a.Equal(1, subarraySum([]int{-1, -1, 1}, 1))
	a.Equal(3, subarraySum([]int{-2, -1, 0, 1, 2}, 0))
}
