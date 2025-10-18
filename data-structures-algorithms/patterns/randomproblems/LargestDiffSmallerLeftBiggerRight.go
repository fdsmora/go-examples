package randomproblems

// my solution
/* func maximumPositiveDifference(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxIdx, minIdx := 1, 0
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
		if nums[j] < nums[minIdx] {
			minIdx = j
		}
		j++
	}
	if (minIdx >= maxIdx) || nums[minIdx] >= nums[maxIdx] {
		return 0
	}
	return nums[maxIdx] - nums[minIdx]
} */

// runtime : O(n)
// space : O(1)
func maximumPositiveDifference(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	minVal := nums[0]
	maxDiff := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-minVal > maxDiff {
			maxDiff = nums[i] - minVal
		}
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return maxDiff
}
