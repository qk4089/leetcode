package _189

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 2)
	a.Equal([]int{5, 6, 1, 2, 3, 4}, nums)
}
