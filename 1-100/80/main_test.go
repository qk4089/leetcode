package _80

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 3, 3}))
}
