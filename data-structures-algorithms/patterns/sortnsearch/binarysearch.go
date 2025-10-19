package sortnsearch

func BinarySearchRecursive(nums []int, start, end, value int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2 // to avoid overflow: start + (end-start)/2
	if nums[mid] == value {
		return mid
	}
	if value > nums[mid] {
		return BinarySearch(nums, mid+1, end, value)
	}
	return BinarySearch(nums, start, mid-1, value)
}

func BinarySearch(nums []int, start, end, value int) int {
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == value {
			return mid
		}
		if value > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
