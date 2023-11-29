package _15

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{-8, 3, 3, 4, 4, 5, 5}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
}
