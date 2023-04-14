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

func Test_beibao(t *testing.T) {
	w, items := 9, []int{2, 2, 4, 6, 3}
	fmt.Println(recall(w, items))
	fmt.Println(dynamic(w, items))
	fmt.Println(dynamic2(w, items))
}

func Test_beibao_with_price(t *testing.T) {
	w, items, values := 9, []int{2, 2, 4, 6, 3}, []int{3, 4, 8, 9, 6}
	fmt.Println(dynamic_with_price(w, items, values))
	//fmt.Println(dynamic_with_price2(w, items, values))
	fmt.Println(dynamic_with_price3(w, items, values))
}
