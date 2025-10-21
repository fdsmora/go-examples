package patterns

import "math"

func getSecondLargestDistinctNumber(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	largest := nums[0]
	largest2 := math.MinInt32
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest2 = largest
			largest = nums[i]
		} else if nums[i] < largest && nums[i] > largest2 {
			largest2 = nums[i]
		}
	}
	if largest2 == math.MinInt32 {
		largest2 = largest
	}

	return largest2
}
