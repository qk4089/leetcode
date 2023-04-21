package _1106

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(parseBoolExpr("!(f)"))
	a.True(parseBoolExpr("|(f,t)"))
	a.False(parseBoolExpr("&(t,f)"))
	a.False(parseBoolExpr("|(&(t,f,t),!(t))"))
}
