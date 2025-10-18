package slidingwindow

import "math"

// leetcode problem 209
func MinSubArrayLen(target int, nums []int) int {
	var s, sum int
	min := math.MaxInt32
	var targetFound bool

	for e := 0; e < len(nums); e++ {
		sum += nums[e]

		for sum >= target {
			min = int(math.Min(float64(e-s+1), float64(min)))
			targetFound = true
			if s == e {
				break
			}
			sum -= nums[s]
			s++
		}
	}
	if !targetFound {
		return 0
	}
	return min
}
