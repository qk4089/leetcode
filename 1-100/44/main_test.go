package _44

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.False(isMatch("aa", "a"))
	a.True(isMatch("bc", "b*?"))
	a.True(isMatch("abba", "*ba"))
	a.True(isMatch("acba", "a*b*a"))
	a.True(isMatch("", "******"))
	a.False(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "***aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}
