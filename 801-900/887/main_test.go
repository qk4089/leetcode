package _887

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(2, superEggDrop(1, 2))
	a.Equal(3, superEggDrop(2, 6))
	a.Equal(4, superEggDrop(3, 14))
}
