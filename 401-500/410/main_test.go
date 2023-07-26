package _410

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(18, splitArray([]int{7, 2, 5, 10, 8}, 2))
	a.Equal(9, splitArray([]int{1, 2, 3, 4, 5}, 2))
	a.Equal(4, splitArray([]int{1, 4, 4}, 3))
	a.Equal(3, splitArray([]int{2, 3, 1, 1, 1, 1}, 5))
}
