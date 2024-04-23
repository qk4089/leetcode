package _33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(3, findMinIdx([]int{3, 4, 5, 1, 2}))
	//a.Equal(1, findMinIdx([]int{5, 1, 2, 3, 4}))
	//a.Equal(-1, search([]int{1}, 2))
	//a.Equal(2, search([]int{5, 1, 2}, 2))
	//a.Equal(1, search([]int{5, 1, 2}, 1))
	//a.Equal(1, search([]int{1, 2}, 2))
	//a.Equal(3, search([]int{3, 4, 5, 1, 2}, 1))
	a.Equal(1, search([]int{5, 1, 2, 3, 4}, 1))
	//a.Equal(1, search([]int{1, 2, 3, 4, 5}, 2))
}
