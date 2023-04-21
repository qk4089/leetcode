package beibao

import (
	"fmt"
	"testing"
)

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
