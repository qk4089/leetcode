package _136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(-1, main([]int{2, -1, 2}))
	a.Equal(4, main([]int{4, 1, 2, 1, 2}))
	a.Equal(1, main([]int{1}))
	a.Equal(-2, main([]int{-1, -1, -2, 1, 1}))

}
