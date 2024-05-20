package _496

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_51(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([]int{-1, 3, -1}, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	a.ElementsMatch([]int{7, 7, 7, 7, 7}, nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
