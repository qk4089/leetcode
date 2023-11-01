package _295

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	a.Equal(1.0, medianFinder.FindMedian())
	medianFinder.AddNum(2)
	a.Equal(1.5, medianFinder.FindMedian())
	medianFinder.AddNum(3)
	a.Equal(2.0, medianFinder.FindMedian())
	medianFinder.AddNum(4)
	a.Equal(2.5, medianFinder.FindMedian())
	medianFinder.AddNum(5)
	a.Equal(3.0, medianFinder.FindMedian())
	medianFinder.AddNum(6)
	a.Equal(3.5, medianFinder.FindMedian())
	medianFinder.AddNum(7)
	a.Equal(4.0, medianFinder.FindMedian())
	medianFinder.AddNum(8)
	a.Equal(4.5, medianFinder.FindMedian())
	medianFinder.AddNum(9)
	a.Equal(5.0, medianFinder.FindMedian())
	medianFinder.AddNum(10)
	a.Equal(5.5, medianFinder.FindMedian())
}
