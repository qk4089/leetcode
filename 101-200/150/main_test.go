package _150

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, evalRPN([]string{"4", "13", "5", "/", "+"}))
}
