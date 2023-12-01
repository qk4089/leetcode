package _29

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(-3, divide(-10, 3))
	a.Equal(-2, divide(7, -3))
	a.Equal(1, divide(2147483640, 2147483630))
	a.Equal(2147483647, divide(-2147483648, -1))
	a.Equal(-2147483648, divide(-2147483648, 1))
	a.Equal(-1073741824, divide(-2147483648, 2))
}
