package _1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(17, maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
