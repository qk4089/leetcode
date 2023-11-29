package _1

//https://leetcode.cn/problems/two-sum/
//给定一个整数数组nums和一个整数目标值target，请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。你可以按任意顺序返回答案。

//提示：
//	2<=nums.length<=10^4
//	-10^9<=nums[i]<=10^9
//	-10^9<=target<=10^9
//	只会存在一个有效答案
//	进阶：你可以想出一个时间复杂度小于O(n^2)的算法吗？

func towSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := index[target-nums[i]]; ok {
			return []int{val, i}
		} else {
			index[nums[i]] = i
		}
	}
	return nil
}
