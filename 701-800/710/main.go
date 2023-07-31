package _710

import (
	"math/rand"
	"sort"
)

//给定一个整数n和一个无重复黑名单整数数组blacklist。设计一种算法，从[0, n - 1]范围内的任意整数中选取一个未加入黑名单blacklist的整数。
//任何在上述范围内且不在黑名单 blacklist 中的整数都应该有同等的可能性被返回。

//优化你的算法，使它最小化调用语言 内置 随机函数的次数。

//实现 Solution 类:
//
//Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
//int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数

//示例 1：
//输入
//["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
//[[7, [2, 3, 5]], [], [], [], [], [], [], []]
//输出
//[null, 0, 4, 1, 6, 1, 0, 4]
//解释
//Solution solution = new Solution(7, [2, 3, 5]);
//solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
//// 0、1、4和6的返回概率必须相等(即概率为1/4)。
//solution.pick(); // 返回 4
//solution.pick(); // 返回 1
//solution.pick(); // 返回 6
//solution.pick(); // 返回 1
//solution.pick(); // 返回 0
//solution.pick(); // 返回 4

//提示:
//	1 <= n <= 10^9
//	0 <= blacklist.length <= min(10^5, n - 1)
//	0 <= blacklist[i] < n
//	blacklist 中所有值都不同
//	pick 最多被调用 2 * 10^4 次

type Solution struct {
	size int
	idx  map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	sort.Ints(blacklist)
	last, sz, idx := n-1, n-len(blacklist), make(map[int]int)
	for _, v := range blacklist {
		idx[v] = 0
	}
	for _, v := range blacklist {
		for _, ok := idx[last]; ok; {
			last--
			_, ok = idx[last]
		}
		idx[v], last = last, last-1
	}
	return Solution{sz, idx}
}

func (this *Solution) Pick() int {
	idx := rand.Intn(this.size)
	if v, ok := this.idx[idx]; ok {
		return v
	}
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
