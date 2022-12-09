package sort

// 假设我们现在需要对 D，a，F，B，c，A，z 这个字符串进行排序，要求将其中所有小写字母都排在大写字母的前面，
// 但小写字母内部和大写字母内部不要求有序，如何来实现呢？如果字符串中存储的不仅有大小写字母，还有数字。
// 要将小写字母的放到前面，大写字母放在最后，数字放在中间，不用排序算法，又该怎么解决呢？
func organize(s string) string {
	start, end, b := 0, len(s)-1, []byte(s)
	for start < end {
		if b[start] < 'a' {
			b[start], b[end] = b[end], b[start]
			end--
		} else {
			start++
		}
	}
	return string(b)
}

// 给定一个整数数组，给定一个值K，这个值在原数组中一定存在，要求把数组中小于K的元素放到数组的左边，大于K的元素放到数组的右边，
// 等于K的元素放到数组的中间，最终返回一个整数数组
// [2, 3, 1, 4, 9, 7, 6, 1, 4, 5] 给定值 K = 4
func organizeWithNum(arr []int, k int) []int {
	point, start, end := 0, 0, len(arr)-1
	for point < end {
		if arr[point] < k {
			arr[start], arr[point] = arr[point], arr[start]
			start++
			point++
		} else if arr[point] > k {
			arr[point], arr[end] = arr[end], arr[point]
			end--
		} else {
			point++
		}
	}
	return arr
}

// 二分查找
func bSearch(arr []int, n int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == n {
			return mid
		} else if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 二分查找-查找等于给定值的第一个元素
func bSearchFirst(arr []int, n int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == n {
			if mid == 0 || arr[mid-1] != n {
				return mid
			} else {
				right = mid - 1
			}
		} else if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 二分查找-查找等于给定值的最后一个元素
func bSearchLast(arr []int, n int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == n {
			if mid == len(arr)-1 || arr[mid+1] != n {
				return mid
			} else {
				left = mid + 1
			}
		} else if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
