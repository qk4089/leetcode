package _1704

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(main("book"))
	a.False(main("textbook"))
}
