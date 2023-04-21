package _322

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, coinChange([]int{1, 2, 5}, 11))
	a.Equal(-1, coinChange([]int{2}, 3))
	a.Equal(0, coinChange([]int{1}, 0))
}
