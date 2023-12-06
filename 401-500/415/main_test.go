package _415

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("134", addStrings("11", "123"))
	a.Equal("10", addStrings("1", "9"))
}
