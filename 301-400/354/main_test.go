package _354

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, maxEnvelopes([][]int{{1, 15}, {7, 18}, {7, 6}, {7, 100}, {2, 200}, {17, 30}, {17, 45}, {3, 5}, {7, 8}, {3, 6}, {3, 10}, {7, 20}, {17, 3}, {17, 45}}))
}
