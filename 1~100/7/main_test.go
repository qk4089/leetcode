package _7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(321, reverse(123))
	a.Equal(-321, reverse(-123))
	a.Equal(21, reverse(120))
	a.Equal(0, reverse(2000000009))
}
