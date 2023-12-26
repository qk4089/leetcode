package _88

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	arr1 := []int{4, 5, 6, 0, 0, 0}
	merge(arr1, 3, []int{1, 9, 11}, 3)
	a.ElementsMatch([]int{1, 4, 5, 6, 9, 11}, arr1)
}
