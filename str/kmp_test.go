package str

import (
	"fmt"
	"testing"
)

func Test_kmp(t *testing.T) {
	fmt.Println(kmp("ababad", "abad"))
}
func Test_next(t *testing.T) {
	fmt.Println(getNext([]rune("abaabab")))
}
