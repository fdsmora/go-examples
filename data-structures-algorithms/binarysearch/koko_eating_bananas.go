package binarysearch

import (
	"math"
	"slices"
)

// LeetCode problem 875. Koko Eating Bananas
// Time complexity: O(n log m) where n is the number of piles and m is the maximum number of bananas in a pile. This is because we perform a binary search on the range of possible eating speeds (from 1 to max bananas in a pile), and for each speed, we calculate the total hours needed to eat all bananas, which takes O(n) time.
// Space complexity: O(1) since we are using a constant amount of extra space.
func minEatingSpeed(piles []int, h int) int {
	max := slices.Max(piles)
	min := 1
	var k, res int
	for min <= max {
		k = (min + max) / 2
		total := sum(k, piles)
		if total <= h {
			res = k
			max = k - 1
		} else {
			min = k + 1
		}
	}
	return res
}

func sum(k int, piles []int) int {
	total := 0.0
	for _, v := range piles {
		total += math.Ceil(float64(v) / float64(k))
	}
	return int(total)
}
