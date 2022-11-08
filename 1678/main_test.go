package _1678

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("Goal", main("G()(al)"))
	a.Equal("Gooooal", main("G()()()()(al)"))
	a.Equal("alGalooG", main("(al)G(al)()()G"))
}
