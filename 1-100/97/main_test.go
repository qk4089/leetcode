package _97

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
