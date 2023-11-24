package _130

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	xx := [][]byte{{'X', 'X', 'X', 'X'}, {'X', '0', '0', 'X'}, {'X', 'X', '0', 'X'}, {'X', '0', 'X', 'X'}}
	solve(xx)
	for i := 0; i < len(xx); i++ {
		fmt.Println(string(xx[i]))
	}
}
