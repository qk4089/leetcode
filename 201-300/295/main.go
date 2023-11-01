package _295

import (
	"container/heap"
)

//中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
//例如 arr = [2,3,4] 的中位数是 3 。
//例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。

//实现 MedianFinder 类:
//	MedianFinder() 初始化 MedianFinder 对象。
//	void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
//	double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。

//示例 1：
//输入
//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//输出
//[null, null, null, 1.5, null, 2.0]
//
//解释
//MedianFinder medianFinder = new MedianFinder();
//medianFinder.addNum(1);    // arr = [1]
//medianFinder.addNum(2);    // arr = [1, 2]
//medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
//medianFinder.addNum(3);    // arr[1, 2, 3]
//medianFinder.findMedian(); // return 2.0

//提示:
//	-10^5 <= num <= 10^5
//	在调用 findMedian 之前，数据结构中至少有一个元素
//	最多 5 * 10^4 次调用 addNum 和 findMedian

type MedianFinder struct {
	small *BigHeap
	big   *SmallHeap
}

func Constructor() MedianFinder {
	var small, big = make(AbsHeap, 0), make(AbsHeap, 0)
	return MedianFinder{&BigHeap{small}, &SmallHeap{big}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.small.Len() >= this.big.Len() {
		heap.Push(this.small, float64(num))
		heap.Push(this.big, heap.Pop(this.small).(float64))
	} else {
		heap.Push(this.big, float64(num))
		heap.Push(this.small, heap.Pop(this.big).(float64))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() == this.big.Len() {
		return (this.small.Peek() + this.big.Peek()) / 2.0
	} else {
		return this.big.Peek()
	}
}

type AbsHeap []float64

func (h AbsHeap) Len() int      { return len(h) }
func (h AbsHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *AbsHeap) Push(x any) { *h = append(*h, x.(float64)) }
func (h *AbsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h AbsHeap) Peek() float64 {
	return h[0]
}

type SmallHeap struct {
	AbsHeap
}

func (h SmallHeap) Less(i, j int) bool { return (h.AbsHeap)[i] < (h.AbsHeap)[j] }

type BigHeap struct {
	AbsHeap
}

func (h BigHeap) Less(i, j int) bool { return (h.AbsHeap)[i] > (h.AbsHeap)[j] }
