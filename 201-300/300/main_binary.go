package _300

func lengthOfLISBinary(nums []int) int {
	idx, result := 0, make([]int, 1)
	result[idx] = nums[0]
	for i := 1; i < len(nums); i++ {
		if result[idx] < nums[i] {
			idx, result = idx+1, append(result, nums[i])
		} else {
			left, right := 0, idx
			for left <= right {
				mid := (left + right) >> 1
				if result[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			result[left] = nums[i]
		}
	}
	return len(result)
}
