package _1704

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(halvesAreAlike("book"))
	a.False(halvesAreAlike("textbook"))
}
