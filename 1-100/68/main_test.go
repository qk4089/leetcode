package _68

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.ElementsMatch([]string{"This    is    an", "example  of text", "justification.  "}, fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	a.ElementsMatch([]string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}, fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}
