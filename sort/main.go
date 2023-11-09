package sort

// 冒泡
func bubbling(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap, arr[j], arr[j+1] = true, arr[j+1], arr[j]
			}
		}
		if !swap {
			return
		}
	}
}

// 插入
func insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, val := i-1, arr[i]
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
}

// 选择
func choice(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

// 归并
func merge(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	p := len(arr) / 2
	return sort(merge(arr[:p]), merge(arr[p:]))
}

func sort(x []int, y []int) []int {
	var result []int
	p0, p1 := 0, 0
	for p0 < len(x) && p1 < len(y) {
		if x[p0] > y[p1] {
			result, p1 = append(result, y[p1]), p1+1
		} else {
			result, p0 = append(result, x[p0]), p0+1
		}
	}
	if p0 == len(x) {
		return append(result, y[p1:]...)
	}
	return append(result, x[p0:]...)
}

// 快排
func fast(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	point, start, pivot := 0, 0, nums[len(nums)-1]
	for point < len(nums)-1 {
		if nums[point] < pivot {
			nums[point], nums[start] = nums[start], nums[point]
			start++
		}
		point++
	}
	nums[point], nums[start] = nums[start], nums[point]
	fast(nums[:start])
	fast(nums[start+1:])
	return nums
}

// 计数
func count(arr []int) {
	if len(arr) == 1 {
		return
	}
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	bucket, result := make([]int, max+1), make([]int, len(arr))
	for _, val := range arr {
		bucket[val] += 1
	}
	for i := 1; i < max+1; i++ {
		bucket[i] = bucket[i] + bucket[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		cur := (arr)[i]
		result[bucket[cur]-1], bucket[cur] = cur, bucket[cur]-1
	}
	for i, val := range result {
		arr[i] = val
	}
}
