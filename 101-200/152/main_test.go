package _152

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(6, maxProduct([]int{2, 3, -2, 4}))
	//a.Equal(0, maxProduct([]int{-2, 0, -1}))
	//a.Equal(24, maxProduct([]int{-2, 3, -4}))
	//a.Equal(2, maxProduct([]int{0, 2}))
	//a.Equal(24, maxProduct([]int{2, -5, -2, -4, 3}))
	a.Equal(108, maxProduct([]int{-1, -2, -9, -6}))
}
