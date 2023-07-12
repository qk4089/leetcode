package _875

//珂珂喜欢吃香蕉。这里n堆香蕉，第i堆中有piles[i]根香蕉。警卫已经离开了，将在h小时后回来。 珂珂可以决定她吃香蕉的速度k（单位：根/小时）。
//每个小时，她将会选择一堆香蕉，从中吃掉 k 根。 如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
//珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。返回她可以在h小时内吃掉所有香蕉的最小速度 k（k 为整数）。

//示例 1：
//输入：piles = [3,6,7,11], h = 8
//输出：4

//输入：piles = [30,11,23,4,20], h = 5
//输出：30

//输入：piles = [30,11,23,4,20], h = 6
//输出：23

//提示：
//	1 <= piles.length <= 10^4
//	piles.length <= h <= 10^9
//	1 <= piles[i] <= 10^9

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for i := 0; i < len(piles); i++ {
		if piles[i] > right {
			right = piles[i]
		}
	}
	if h == len(piles) {
		return right
	}
	for left < right {
		mid := left + (right-left)>>1
		if eat(mid, h, piles) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func eat(val, target int, piles []int) bool {
	count := 0
	for i := 0; i < len(piles); i++ {
		count += piles[i] / val
		if piles[i]%val != 0 {
			count++
		}
	}
	return count <= target
}
