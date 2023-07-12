package _35

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, searchInsert([]int{1, 3, 5, 6}, 2))
}
