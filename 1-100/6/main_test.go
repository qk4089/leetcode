package _6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal("PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	a.Equal("AB", convert("AB", 1))
}
