package _1678

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("Goal", interpret("G()(al)"))
	a.Equal("Gooooal", interpret("G()()()()(al)"))
	a.Equal("alGalooG", interpret("(al)G(al)()()G"))
}
