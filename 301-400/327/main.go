package _327

import "sort"

//给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 区间和的个数。
//区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

//示例 1：
//输入：nums = [-2,5,-1], lower = -2, upper = 2
//输出：3
//解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。

//输入：nums = [0], lower = 0, upper = 0
//输出：1

//提示：
//	1 <= nums.length <= 10^5
//	-2^31 <= nums[i] <= 2^31 - 1
//	-10^5 <= lower <= upper <= 10^5
//题目数据保证答案是一个 32 位 的整数

type BitTree struct {
	tree []int
}

func (bt BitTree) add(i int) {
	for ; i < len(bt.tree); i += i & -i {
		bt.tree[i]++
	}
}

func (bt BitTree) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += bt.tree[i]
	}
	return
}

func (bt BitTree) query(l, r int) (res int) {
	return bt.sum(r) - bt.sum(l-1)
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	ans, preSum, allSum := 0, make([]int, len(nums)+1), make([]int, 3*len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		allSum = append(allSum, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}
	sort.Ints(allSum)
	idx, idxMap := 1, map[int]int{allSum[0]: 1}
	for i := 1; i < len(allSum); i++ {
		if allSum[i] != allSum[i-1] {
			idx++
			idxMap[allSum[i]] = idx
		}
	}
	tree := BitTree{make([]int, idx+1)}
	tree.add(idxMap[0])
	for i := 1; i < len(preSum); i++ {
		ans += tree.query(idxMap[preSum[i]-upper], idxMap[preSum[i]-lower])
		tree.add(idxMap[preSum[i]])
	}
	return ans

	//for i := 1; i < len(preSum); i++ {
	//	for j := 0; j < i; j++ {
	//		sum := preSum[i] - preSum[j]
	//		if sum >= lower && sum <= upper {
	//			ans++
	//		}
	//	}
	//}
	//return ans
}
