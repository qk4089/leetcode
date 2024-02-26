package _131

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
}
