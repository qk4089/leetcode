package _239

import "container/heap"

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。 返回滑动窗口中的最大值 。

//示例 1：
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7

//输入：nums = [1], k = 1
//输出：[1]

//提示：
//	1 <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4
//	1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {
	//return useHeap(nums, k)
	return useQueue(nums, k)
}

func useQueue(nums []int, k int) []int {
	dqueue := &DQueue{nums, []int{}}
	for i := 0; i < k; i++ {
		dqueue.Push(i)
	}
	ans := []int{dqueue.Top()}
	for i := k; i < len(nums); i++ {
		for dqueue.Len() > 0 && dqueue.idx[0] <= i-k {
			dqueue.Pop()
		}
		dqueue.Push(i)
		ans = append(ans, dqueue.Top())
	}
	return ans
}

type DQueue struct {
	nums []int
	idx  []int
}

func (this DQueue) Len() int {
	return len(this.idx)
}
func (this *DQueue) Push(val int) {
	if len(this.idx) == 0 {
		this.idx = []int{val}
		return
	}
	for i := 0; i < len(this.idx); i++ {
		if this.nums[val] >= this.nums[this.idx[i]] {
			this.idx = this.idx[:i]
			break
		}
	}
	this.idx = append(this.idx, val)
}

func (this *DQueue) Pop() {
	this.idx = this.idx[1:]
}

func (this DQueue) Top() int {
	return this.nums[this.idx[0]]
}

func useHeap(nums []int, k int) []int {
	h := &BigHeap{nums, []int{}}
	for i := 0; i < k; i++ {
		h.Push(i)
	}
	heap.Init(h)
	ans := []int{h.Top()}
	for i := k; i < len(nums); i++ {
		for h.Len() > 0 && h.idx[0] <= i-k {
			heap.Pop(h)
		}
		heap.Push(h, i)
		ans = append(ans, h.Top())
	}
	return ans
}

type BigHeap struct {
	nums []int
	idx  []int
}

func (h BigHeap) Len() int           { return len(h.idx) }
func (h BigHeap) Less(i, j int) bool { return h.nums[h.idx[i]] > h.nums[h.idx[j]] }
func (h BigHeap) Swap(i, j int)      { h.idx[i], h.idx[j] = h.idx[j], h.idx[i] }

func (h *BigHeap) Push(x any) { h.idx = append(h.idx, x.(int)) }

func (h *BigHeap) Pop() any {
	l := len(h.idx)
	val := h.nums[h.idx[l-1]]
	h.idx = h.idx[:l-1]
	return val
}

func (h BigHeap) Top() int {
	return h.nums[h.idx[0]]
}
