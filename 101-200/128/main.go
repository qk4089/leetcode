package _128

//https://leetcode.cn/problems/longest-consecutive-sequence/
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

//示例 1：
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9

//提示：
//	0 <= nums.length <= 10^5
//	-10^9 <= nums[i] <= 10^9

func longestConsecutive(nums []int) int {
	maxCnt, cntMap := 0, make(map[int]int)
	for _, num := range nums {
		if _, ok := cntMap[num]; !ok {
			cntMap[num]++
		}
	}
	for key, val := range cntMap {
		if _, ok := cntMap[key-1]; ok {
			continue
		}
		for _, ok := cntMap[key+1]; ok; _, ok = cntMap[key+1] {
			val++
			key++
		}
		maxCnt = max(maxCnt, val)
	}
	return maxCnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
