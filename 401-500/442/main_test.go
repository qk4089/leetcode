package _442

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{2, 3}, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	a.Equal([]int{1}, findDuplicates([]int{1, 1, 2}))
	a.Equal([]int{}, findDuplicates([]int{1}))
	a.Equal([]int{2}, findDuplicates([]int{2, 2}))
}
