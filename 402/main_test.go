package _402

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("1219", removeKdigits("1432219", 3))
	a.Equal("200", removeKdigits("10200", 1))
	a.Equal("0", removeKdigits("10001", 4))
	a.Equal("11", removeKdigits("112", 1))
	a.Equal("1", removeKdigits("10001", 1))
}
