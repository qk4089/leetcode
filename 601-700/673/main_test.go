package _673

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, findNumberOfLIS([]int{1, 3, 4, 9, 2}))
	a.Equal(2, findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	a.Equal(5, findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	a.Equal(1, findNumberOfLIS([]int{1}))
	a.Equal(3, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
	a.Equal(2, findNumberOfLIS([]int{1, 3, 2, 4, 5}))
	a.Equal(1, findNumberOfLIS([]int{2, 3, 1}))
}
