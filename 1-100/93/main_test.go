package _93

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([]string{"0.0.0.0"}, restoreIpAddresses("0000"))
	a.ElementsMatch([]string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))
	a.ElementsMatch([]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses("101023"))
}
