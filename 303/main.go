package _303

//给定一个整数数组nums，处理以下类型的多个查询: 计算索引left和right（包含left和right）之间的nums元素的和，其中left<=right
//实现NumArray类：
//	NumArray(int[]nums)使用数组nums初始化对象
//	int sumRange(int i,int j)返回数组nums中索引left和right之间的元素的总和，包含left和right两点（也就是nums[left]+nums[left+1]+...+nums[right])

//提示：
//	1<=nums.length<=104
//	-105<=nums[i]<=105
//	0<=i<=j<nums.length
//	最多调用104次sumRange方法

//示例1
//输入：
//	["NumArray","sumRange","sumRange","sumRange"]
//	[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]
//输出：
//	[null,1,-1,-3]
//解释：
//	NumArray numArray=newNumArray([-2,0,3,-5,2,-1]);
//	numArray.sumRange(0,2);//return 1 ((-2)+0+3)
//	numArray.sumRange(2,5);//return -1 (3+(-5)+2+(-1))
//	numArray.sumRange(0,5);//return -3 ((-2)+0+3+(-5)+2+(-1))

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	result := make([]int, len(nums)+1)
	result[0] = 0
	for i, num := range nums {
		result[i+1] = result[i] + num
	}
	return NumArray{sum: result}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
