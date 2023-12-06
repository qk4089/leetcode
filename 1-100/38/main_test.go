package _38

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("21", countAndSay(3))
	a.Equal("1211", countAndSay(4))
	a.Equal("111221", countAndSay(5))
	a.Equal("312211", countAndSay(6))
}
