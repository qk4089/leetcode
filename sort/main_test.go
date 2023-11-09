package sort

import (
	"fmt"
	"testing"
)

func Test_bubbling(t *testing.T) {
	arr := []int{3, 6, 1, 5, 4, 2}
	bubbling(arr)
	fmt.Println(arr)
}

func Test_insert(t *testing.T) {
	arr := []int{3, 6, 1, 5, 4, 2}
	insert(arr)
	fmt.Println(arr)
}

func Test_choice(t *testing.T) {
	arr := []int{3, 6, 1, 5, 4, 2}
	choice(arr)
	fmt.Println(arr)
}

func Test_merge(t *testing.T) {
	arr := []int{3, 6, 1, 5, 4, 2}
	merge(arr)
	fmt.Println(arr)
}

func Test_fast(t *testing.T) {
	fmt.Println(fast([]int{3, 6, 1, 5, 4, 2}))
}

func Test_count(t *testing.T) {
	arr := []int{2, 5, 3, 0, 2, 3, 0, 3}
	count(arr)
	fmt.Println(arr)
}
