package _384

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("accaccacc", decodeString("3[a2[c]]"))
}
