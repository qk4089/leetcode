package _162

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, findPeakElement([]int{1, 2, 3, 1}))
	a.Equal(1, findPeakElement([]int{1, 6, 5, 4, 3, 2, 1}))
}
