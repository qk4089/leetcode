package _355

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	a.ElementsMatch([]int{5}, twitter.GetNewsFeed(1))
}
