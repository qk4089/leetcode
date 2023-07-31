package _710

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := Constructor(7, []int{2, 3, 5})
	fmt.Println(s.Pick())
}
