package _8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(42, myAtoi("42"))
	//a.Equal(4193, myAtoi("4193 with words"))
	//a.Equal(-2147483647, myAtoi("-2147483647"))
	a.Equal(2147483647, myAtoi("20000000000000000000"))
	//a.Equal(2147483647, myAtoi("21474836460"))
	a.Equal(12345678, myAtoi("  0000000000012345678"))
	//a.Equal(-2147483648, myAtoi("-6147483648"))
	//a.Equal(2147483647, myAtoi("2147483648"))
}
