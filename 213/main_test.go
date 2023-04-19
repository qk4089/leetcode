package _213

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(3, rob([]int{2, 3, 2}))
	//a.Equal(3, rob([]int{1, 2, 1, 1}))
	//a.Equal(3, rob([]int{1, 2, 3}))
	//a.Equal(2, rob([]int{1, 2}))
	a.Equal(340, rob([]int{200, 3, 140, 20, 10}))
}
