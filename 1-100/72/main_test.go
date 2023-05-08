package _72

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDistance(t *testing.T) {
	a := assert.New(t)
	//a.Equal(3, recall("horse", "ros"))
	//a.Equal(5, recall("intention", "execution"))
	a.Equal(3, minDistance("horse", "ros"))
	a.Equal(5, minDistance("intention", "execution"))
	a.Equal(1, minDistance("abc", "ab"))
}
