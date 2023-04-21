package _167

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{1, 2}, main([]int{2, 7, 11, 16}, 9))
	a.Equal([]int{1, 2}, main([]int{-1, 0}, -1))
}
