package offer

import "sort"

//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

//示例 1:
//输入: [7,5,6,4]
//输出: 5

//限制：
//0 <= 数组长度 <= 50000

func reversePairs(nums []int) int {
	return merge(nums)
}

func merge(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	sum, p0, p1, tmp := merge(nums[:mid])+merge(nums[mid:]), 0, mid, make([]int, 0)
	for p0 < mid && p1 < len(nums) {
		if nums[p0] <= nums[p1] {
			tmp, p0 = append(tmp, nums[p0]), p0+1
			sum += p1 - mid
		} else {
			tmp, p1 = append(tmp, nums[p1]), p1+1
		}
	}
	if p0 == mid {
		tmp = append(tmp, nums[p1:]...)
	}
	if p1 == len(nums) {
		tmp = append(tmp, nums[p0:]...)
		sum += (len(nums) - mid) * (mid - p0)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
	return sum
}

// 使用树状数组
func reversePairs2(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < len(nums); i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}
	cnt, bit := 0, make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i] - 1
		for x > 0 {
			cnt += bit[x]
			x -= x & (-x)
		}
		for idx := nums[i]; idx < len(bit); idx += idx & (-idx) {
			bit[idx] += 1
		}
	}
	return cnt
}
