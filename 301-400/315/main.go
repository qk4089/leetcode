package _315

//给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

//示例 1：
//输入：nums = [5,2,6,1]
//输出：[2,1,1,0]
//解释：
//5 的右侧有 2 个更小的元素 (2 和 1)
//2 的右侧仅有 1 个更小的元素 (1)
//6 的右侧有 1 个更小的元素 (1)
//1 的右侧有 0 个更小的元素

//输入：nums = [-1]
//输出：[0]

//输入：nums = [-1,-1]
//输出：[0,0]

//提示：
//	1 <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4

func countSmaller(nums []int) []int {
	tuples, result := make([]Tuple, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tuples[i] = Tuple{val: nums[i], idx: i}
	}
	merge(tuples, &result)
	return result
}

func merge(nums []Tuple, ans *[]int) []Tuple {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	x, y := merge(nums[:mid], ans), merge(nums[mid:], ans)
	p0, p1, result := 0, 0, make([]Tuple, 0)
	for p0 < len(x) && p1 < len(y) {
		if x[p0].val <= y[p1].val {
			(*ans)[x[p0].idx] += p1
			p0, result = p0+1, append(result, x[p0])
		} else {
			p1, result = p1+1, append(result, y[p1])
		}
	}
	if p0 == len(x) {
		return append(result, y[p1:]...)
	}
	for i := p0; i < len(x); i++ {
		(*ans)[x[i].idx] += len(y)
	}
	return append(result, x[p0:]...)
}

type Tuple struct {
	val int
	idx int
}
