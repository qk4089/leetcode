package _71

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("/home", simplifyPath("/home/"))
	a.Equal("/s", simplifyPath("s"))
	a.Equal("/home/foo", simplifyPath("/home//foo/"))
	a.Equal("/", simplifyPath("/../"))
	a.Equal("/asd", simplifyPath("/asd/."))
	a.Equal("/asd.a", simplifyPath("/asd.a"))
	a.Equal("/asd/.a", simplifyPath("/asd/.a"))
	a.Equal("/", simplifyPath("/asd/.."))
	a.Equal("/asd", simplifyPath("/asd/."))
}
