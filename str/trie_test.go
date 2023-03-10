package str

import (
	"fmt"
	"testing"
)

func Test_Tire(t *testing.T) {
	tire := NewTire()
	tire.insert("hello")
	tire.insert("he")
	tire.insert("sherry")
	fmt.Println(tire.find("she"))
}
