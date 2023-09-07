package misc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_organize(t *testing.T) {
	fmt.Println(organize("DaFBcAz"))
}

func Test_organize_with_num(t *testing.T) {
	fmt.Println(organizeWithNum([]int{2, 3, 1, 4, 9, 7, 6, 1, 4, 5}, 4))
}

func Test_b_search(t *testing.T) {
	a := assert.New(t)
	arr := []int{2, 3, 6, 8, 9, 11, 15, 19}
	a.Equal(1, bSearch(arr, 3))
	a.Equal(-1, bSearch(arr, 5))
}

func Test_b_search_first(t *testing.T) {
	a := assert.New(t)
	a.Equal(-1, bSearchFirst([]int{3, 3, 3, 8, 9, 11, 15, 19}, 5))
	a.Equal(0, bSearchFirst([]int{3, 3, 3, 8, 9, 11, 15, 19}, 3))
	a.Equal(1, bSearchFirst([]int{1, 3, 3, 3, 8, 9, 11, 15, 19}, 3))
}

func Test_b_search_last(t *testing.T) {
	a := assert.New(t)
	a.Equal(-1, bSearchLast([]int{3, 3, 3, 8, 9, 11, 15, 19}, 5))
	a.Equal(7, bSearchLast([]int{1, 3, 3, 3, 8, 9, 11, 15, 19}, 15))
	a.Equal(8, bSearchLast([]int{1, 3, 3, 3, 8, 9, 11, 15, 15}, 15))
}

func Test_findQueens(t *testing.T) {
	findQueens(0, make([]int, 8))
}

func Test_flipCount(t *testing.T) {
	a := assert.New(t)
	a.EqualValues(two(3, 3), two(flipCount("BBFBFBB")))
	a.EqualValues(two(2, 4), two(flipCount("BBFBFBBB")))
}

func two(a, b any) []any {
	return []any{a, b}
}
