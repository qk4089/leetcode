package _41

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, firstMissingPositive([]int{1, 2, 0}))
	a.Equal(2, firstMissingPositive([]int{3, 4, -1, 1}))
	a.Equal(2, firstMissingPositive([]int{3, 1}))
	a.Equal(5, firstMissingPositive([]int{1, 2, 3, 4}))
	a.Equal(1, firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
