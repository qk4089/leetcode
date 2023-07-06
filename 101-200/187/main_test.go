package _187

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAA"))
}
