package _91

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, numDecodings("27"))
	a.Equal(1, numDecodings("2101"))
	a.Equal(1, numDecodings("220"))
	a.Equal(0, numDecodings("2200"))
	a.Equal(0, numDecodings("230"))
	a.Equal(1, numDecodings("101"))
	a.Equal(0, numDecodings("10021"))
	a.Equal(1, numDecodings("10"))
}
