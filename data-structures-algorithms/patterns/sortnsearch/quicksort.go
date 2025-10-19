package sortnsearch

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	qSort(nums, 0, len(nums)-1)
}

func qSort(nums []int, left, right int) {
	index := partition(nums, left, right)
	if left < index-1 {
		qSort(nums, left, index-1)
	}
	if index < right {
		qSort(nums, index, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left+(right-left)/2]

	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		if left <= right {
			swap(nums, left, right)
			left++
			right--
		}
	}
	return left
}

func swap(nums []int, left, right int) {
	tmp := nums[left]
	nums[left] = nums[right]
	nums[right] = tmp
}
