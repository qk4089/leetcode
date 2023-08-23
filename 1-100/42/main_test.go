package _42

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	a.Equal(9, trap([]int{4, 2, 0, 3, 2, 5}))
	a.Equal(0, trap([]int{1}))
}
