package _27

//给你一个数组nums和一个值val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用O(1)额外空间并原地修改输入数组。 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

//提示：
//	0<=nums.length<=100
//	0<=nums[i]<=50
//	0<=val<=100

//示例
//输入：nums=[3,2,2,3],val=3
//输出：2,nums=[2,2]
//解释：函数应该返回新的长度2,并且nums中的前两个元素均为2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为2，而nums=[2,2,3,3]或nums=[2,2,0,0]，也会被视作正确答案。

//输入：nums=[0,1,2,2,3,0,4,2],val=2
//输出：5,nums=[0,1,4,0,3]
//解释：函数应该返回新的长度5,并且nums中的前五个元素为0,1,3,0,4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
