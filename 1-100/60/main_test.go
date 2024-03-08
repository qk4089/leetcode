package _60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("132", getPermutation(3, 2))
	a.Equal("213", getPermutation(3, 3))
	a.Equal("1432", getPermutation(4, 6))
	a.Equal("2314", getPermutation(4, 9))
}
