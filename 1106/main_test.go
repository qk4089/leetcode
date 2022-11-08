package _1106

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(main("!(f)"))
	a.True(main("|(f,t)"))
	a.False(main("&(t,f)"))
	a.False(main("|(&(t,f,t),!(t))"))
}
