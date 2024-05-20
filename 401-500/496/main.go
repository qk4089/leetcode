package _496

//https://leetcode.cn/problems/next-greater-element-i/
//nums1中数字x的下一个更大元素是指x在nums2中对应位置右侧的第一个比x大的元素。给你两个没有重复元素的数组nums1和nums2，
//下标从0开始计数，其中nums1是nums2的子集。对于每个 0 <= i < nums1.length ，找出满足nums1[i] == nums2[j]的下标j，
//并且在nums2确定nums2[j]的下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

//返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

//示例
//输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出：[-1,3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
//- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

//输入：nums1 = [2,4], nums2 = [1,2,3,4].
//输出：[3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
//- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。

//提示：
//	1 <= nums1.length <= nums2.length <= 1000
//	0 <= nums1[i], nums2[i] <= 10^4
//	nums1和nums2中所有整数 互不相同
//	nums1 中的所有整数同样出现在 nums2 中

//进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans, stack, memo := make([]int, len(nums1)), []int{nums2[len(nums2)-1]}, make(map[int]int)
	for i := len(nums2) - 2; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			memo[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		if val, exist := memo[nums1[i]]; exist {
			ans[i] = val
		} else {
			ans[i] = -1
		}
	}
	return ans
}
