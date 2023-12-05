package _45

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(2, jump([]int{2, 3, 1, 1, 4}))
	//a.Equal(2, jump([]int{2, 3, 0, 1, 4}))
	//a.Equal(0, jump([]int{1}))
	//a.Equal(3, jump([]int{1, 2, 1, 1, 1}))
	a.Equal(2, jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
}
