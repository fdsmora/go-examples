package binarysearch

// LC 153.
func findMin(nums []int) int {
	s, e, m := 0, len(nums)-1, 0
	miin := nums[s]

	for s <= e {
		if nums[s] < nums[e] {
			return min(miin, nums[s])
		}

		m = (s + e) / 2
		if nums[m] < miin {
			miin = nums[m]
		}
		if nums[s] <= nums[m] {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return miin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
