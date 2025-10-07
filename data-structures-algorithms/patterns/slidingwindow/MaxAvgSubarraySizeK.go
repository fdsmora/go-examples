package slidingwindow

import "math"

func MaxAvgSubarraySizeK(nums []int, k int) float64 {
	max := math.Inf(-1)
	var windowSum int
	var start int

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]
		if end-start+1 == k {
			max = math.Max(max, float64(windowSum)/float64(k))
			windowSum -= nums[start]
			start++
		}
	}
	return max
}

// another solution
/* func MaxAvgSubarraySizeK(arr []int, k int) (float64, error) {
	if k < len(arr) {
		return 0, errors.New("k < len(arr)")
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	maxSum := sum
	for i := k; i < len(arr); i++ {
		sum += arr[i] + arr[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k), nil
} */
