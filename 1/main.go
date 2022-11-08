package _1

//给定一个整数数组nums和一个整数目标值target，请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。你可以按任意顺序返回答案。

//提示：
//	2<=nums.length<=104
//	-109<=nums[i]<=109
//	-109<=target<=109
//	只会存在一个有效答案
//	进阶：你可以想出一个时间复杂度小于O(n2)的算法吗？

func main(nums []int, target int) []int {
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
