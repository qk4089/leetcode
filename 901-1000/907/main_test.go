package _907

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(17, sumSubarrayMins([]int{3, 1, 2, 4}))
	a.Equal(444, sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
