package _630

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, scheduleCourse([][]int{{1, 2}, {2, 3}}))
	a.Equal(0, scheduleCourse([][]int{{3, 2}, {4, 3}}))
	a.Equal(2, scheduleCourse([][]int{{5, 5}, {2, 6}, {4, 6}}))
	a.Equal(3, scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
}
