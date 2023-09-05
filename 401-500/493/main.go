package _493

//给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
//你需要返回给定数组中的重要翻转对的数量。

//示例 1:
//输入: [1,3,2,3,1]
//输出: 2

//输入: [2,4,3,5,1]
//输出: 3

//注意:
//	给定数组的长度不会超过50000。
//	输入数组中的所有数字都在32位整数的表示范围内。

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	p0, p1, tmp, sum := 0, mid, make([]int, 0), reversePairs(nums[:mid])+reversePairs(nums[mid:])
	for p0 < mid && p1 < len(nums) {
		if nums[p0] < nums[p1] {
			sum += find(nums[mid:], nums[p0])
			p0, tmp = p0+1, append(tmp, nums[p0])
		} else {
			p1, tmp = p1+1, append(tmp, nums[p1])
		}
	}
	for i := p0; i < mid; i++ {
		sum, tmp = sum+find(nums[mid:], nums[i]), append(tmp, nums[i])
	}
	for i := p1; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
	return sum
}

func find(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid]*2 < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if nums[left]*2 < target {
		return left + 1
	} else {
		return left
	}
}
