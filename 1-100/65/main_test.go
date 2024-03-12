package _65

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.False(isNumber("++5"))
	//a.False(isNumber("inf"))
	//a.False(isNumber("."))
	//a.True(isNumber("1"))
	//a.False(isNumber("1e"))
	//a.False(isNumber("e1"))
}
