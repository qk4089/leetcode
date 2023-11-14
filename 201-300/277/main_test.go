package _277

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, findCelebrity([][]int{{1, 1, 0}, {0, 1, 0}, {1, 1, 1}}))
}
