package _354

import "sort"

//https://leetcode.cn/problems/russian-doll-envelopes/
//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。注意：不允许旋转信封。

//示例 1：
//输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出：3
//解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

//输入：envelopes = [[1,1],[1,1],[1,1]]
//输出：1

//提示：
//	1 <= envelopes.length <= 10^5
//	envelopes[i].length == 2
//	1 <= wi, hi <= 10^5

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	stack := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		if len(stack) == 0 || envelopes[i][1] > stack[len(stack)-1] {
			stack = append(stack, envelopes[i][1])
		} else if envelopes[i][1] < stack[0] {
			stack[0] = envelopes[i][1]
		} else {
			left, right := 0, len(stack)-1
			for left < right {
				mid := left + (right-left)/2
				if stack[mid] == envelopes[i][1] {
					break
				} else if stack[mid] > envelopes[i][1] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			if stack[left] > envelopes[i][1] {
				stack[left] = envelopes[i][1]
			}
		}
	}
	return len(stack)
}
