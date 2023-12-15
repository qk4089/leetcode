package _78

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

//示例 1：
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

//输入：nums = [0]
//输出：[[],[0]]

//提示：
//	1 <= nums.length <= 10
//	-10 <= nums[i] <= 10
//	nums 中的所有元素 互不相同

func subsets(nums []int) [][]int {
	ans := [][]int{{}, nums}
	for i := 1; i < len(nums); i++ {
		backtrace(0, i, []int{}, nums, &ans)
	}
	return ans
}

func backtrace(start, cnt int, trace, nums []int, ans *[][]int) {
	if len(trace) == cnt {
		*ans = append(*ans, append([]int{}, trace...))
		return
	}
	for i := start; i < len(nums); i++ {
		trace = append(trace, nums[i])
		backtrace(i+1, cnt, trace, nums, ans)
		trace = trace[:len(trace)-1]
	}
}
