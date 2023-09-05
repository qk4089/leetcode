package _912

//给你一个整数数组 nums，请你将该数组升序排列。
//
//示例 1：
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]

//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]

//提示：
//	1 <= nums.length <= 5 * 10^4
//	-5 * 10^4 <= nums[i] <= 5 * 10^4

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	x, y := sortArray(nums[:mid]), sortArray(nums[mid:])
	p0, p1, ans := 0, 0, make([]int, 0)
	for p0 < len(x) && p1 < len(y) {
		if x[p0] <= y[p1] {
			p0, ans = p0+1, append(ans, x[p0])
		} else {
			p1, ans = p1+1, append(ans, y[p1])
		}
	}
	if p0 == len(x) {
		return append(ans, y[p1:]...)
	}
	return append(ans, x[p0:]...)
}
