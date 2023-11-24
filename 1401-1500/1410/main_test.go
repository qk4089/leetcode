package _1410

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("&>", entityParser("&&gt;"))
	a.Equal("& is an HTML entity but &ambassador; is not.", entityParser("&amp; is an HTML entity but &ambassador; is not."))
	a.Equal("and I quote: \"...\"", entityParser("and I quote: &quot;...&quot;"))
	a.Equal("Stay home! Practice on Leetcode :)", entityParser("Stay home! Practice on Leetcode :)"))
	a.Equal("x > y && x < y is always false", entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
}
