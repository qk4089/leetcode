package _1143

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDistance(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, longestCommonSubsequence("abcde", "ace"))
}
