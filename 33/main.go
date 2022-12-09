package _33

//整数数组nums按升序排列，数组中的值互不相同。在传递给函数之前，nums在预先未知的某个下标k（0<=k<nums.length）上进行了旋转，
//使数组变为[nums[k],nums[k+1],...,nums[n-1],nums[0],nums[1],...,nums[k-1]]（下标从0开始计数）。
//例如，[0,1,2,4,5,6,7]在下标3处经旋转后可能变为[4,5,6,7,0,1,2]。
//给你旋转后的数组nums和一个整数target，如果nums中存在这个目标值target，则返回它的下标，否则返回-1。

//你必须设计一个时间复杂度为O(logn)的算法解决此问题。

//示例
//输入：nums=[4,5,6,7,0,1,2],target=0
//输出：4

//输入：nums=[4,5,6,7,0,1,2],target=3
//输出：-1

//输入：nums=[1],target=0
//输出：-1

//提示：
//	1<=nums.length<=5000
//	-10^4<=nums[i]<=10^4
//	nums中的每个值都独一无二，题目数据保证nums在预先未知的某个下标上进行了旋转
//	-104<=target<=104

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	i := findMin(nums, left, right)
	if target == nums[i] {
		return i
	} else if target < nums[i] {
		return -1
	} else {
		if target == nums[right] {
			return right
		} else if target < nums[right] {
			left = i
		} else if i == 0 {
			return -1
		} else {
			right = i - 1
		}
	}
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findMin(nums []int, left, right int) int {
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
