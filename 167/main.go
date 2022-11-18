package _167

//给你一个下标从1开始的整数数组numbers，该数组已按递增顺序排列，请你从数组中找出满足相加之和等于目标数target的两个数。
//如果设这两个数分别是numbers[index1]和numbers[index2]，则1<=index1<index2<=numbers.length。
//以长度为2的整数数组[index1,index2]的形式返回这两个整数的下标index1和index2。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。你所设计的解决方案必须只使用常量级的额外空间。

//提示：
//	2<=numbers.length<=3*104
//	-1000<=numbers[i]<=1000
//	-1000<=target<=1000
//	仅存在一个有效答案

//示例
//输入：numbers=[2,7,11,15],target=9
//输出：[1,2]
//解释：2与7之和等于目标数9。因此index1=1,index2=2。返回[1,2]。

//输入：numbers=[2,3,4],target=6
//输出：[1,3]
//解释：2与4之和等于目标数6。因此index1=1,index2=3。返回[1,3]。

//输入：numbers=[-1,0],target=-1
//输出：[1,2]
//解释：-1与0之和等于目标数-1。因此index1=1,index2=2。返回[1,2]。

func towSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
}
