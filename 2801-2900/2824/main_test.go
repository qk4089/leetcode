package _2824

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, countPairs([]int{-5, 0, -7, -1, 9, 8, -9, 9}, -14))
}
