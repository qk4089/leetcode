package _341

import (
	"errors"
)

//https://leetcode.cn/problems/flatten-nested-list-iterator/description/
//给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。
//请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。
//实现扁平迭代器类 NestedIterator ：
//	NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
//	int next() 返回嵌套列表的下一个整数。
//	boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。

//你的代码将会用下述伪代码检测：
//	initialize iterator with nestedList
//	res = []
//	while iterator.hasNext()
//		append iterator.next() to the end of res
//	return res
//如果 res 与预期的扁平化列表匹配，那么你的代码将会被判为正确。

//示例 1：
//输入：nestedList = [[1,1],2,[1,1]]
//输出：[1,1,2,1,1]
//解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。

//输入：nestedList = [1,[4,[6]]]
//输出：[1,4,6]
//解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。

//提示：
//	1 <= nestedList.length <= 500
//	嵌套列表中的整数值在范围 [-10^6, 10^6] 内

type NestedIterator struct {
	origin []*NestedInteger
	result int
	size   int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{nestedList, 0, 0}
}

func (this *NestedIterator) Next() int {
	return this.value().(int)
}

func (this *NestedIterator) value() any {
	if this.size > 0 {
		this.size--
		return this.result
	}
	return errors.New("error")
}

func (this *NestedIterator) HasNext() bool {
	if this.size == 0 && len(this.origin) > 0 {
		if val, ok := getValue(&this.origin); ok {
			this.result = val
			this.size++
		}
	}
	return this.size > 0
}

func getValue(elements *[]*NestedInteger) (int, bool) {
	ans, exist := 0, false
	for i := 0; i < len(*elements); i++ {
		element := (*elements)[i]
		if element.IsInteger() {
			ans, exist, *elements = element.GetInteger(), true, (*elements)[i+1:]
			break
		} else {
			item := element.GetList()
			if val, ok := getValue(&item); ok {
				ans, exist = val, ok
				*elements = append(item, (*elements)[i+1:]...)
				break
			}
		}
	}
	return ans, exist
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val  int
	size int
	list []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.size > 0
}
func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetInteger() int {
	this.size--
	return this.val
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}
