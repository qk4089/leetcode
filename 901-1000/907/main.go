package _907

//https://leetcode.cn/problems/sum-of-subarray-minimums/
//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//由于答案可能很大，因此 返回答案模 10^9 + 7 。

//示例 1：
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。

//输入：arr = [11,81,94,43,3]
//输出：444

//提示：
//	1 <= arr.length <= 3 * 10^4
//	1 <= arr[i] <= 3 * 10^4

func sumSubarrayMins(arr []int) int {
	sum, stack := 0, []int{0}
	computed := func(idx, val int) {
		last := len(stack) - 1
		for ; last >= 0 && arr[stack[last]] > val; last-- {
			product := stack[last] + 1
			if last > 0 {
				product = stack[last] - stack[last-1]
			}
			sum += arr[stack[last]] * (idx - stack[last]) * product
		}
		stack = append(stack[:last+1], idx)
	}
	for i := 1; i < len(arr); i++ {
		if arr[stack[len(stack)-1]] < arr[i] {
			stack = append(stack, i)
			continue
		}
		computed(i, arr[i])
	}
	computed(len(arr), 0)
	return sum % (1e9 + 7)
}
