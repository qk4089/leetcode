package str

import (
	"fmt"
	"testing"
)

func Test_bm(t *testing.T) {
	fmt.Println(bm("DaFBcAz", "abc"))
}

func Test_generateGS(t *testing.T) {
	str := "cabcab"
	generateGS(str, make([]int, len(str)), make([]bool, len(str)))
}
