package _34

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{0, 0}, searchRange([]int{1}, 1))
	a.Equal([]int{0, 1}, searchRange([]int{2, 2}, 2))
	a.Equal([]int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
