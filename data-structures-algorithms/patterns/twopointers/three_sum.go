package miscelaneous

import (
	"sort"
)

// LC 15. 3Sum

// LC 15. 3Sum
// My solution, 2 pointers O(n^2) runtime, O(1) space
func threeSum(nums []int) [][]int {
	res := map[[3]int]struct{}{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		k := len(nums) - 1
		j := i + 1
		for j < k {
			k_v := -nums[i] - nums[j]
			for j < k && nums[k] > k_v {
				k--
			}
			if j < k && nums[k] == k_v {
				res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
			}
			j++
		}
	}
	result := [][]int{}
	for triplet := range res {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}
	return result
}

// O(n^2logn)
/* func threeSum(nums []int) [][]int {
	res := map[[3]int]struct{}{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			k_v := -nums[i] - nums[j]
			if k, ok := search(nums, j+1, len(nums)-1, k_v); ok {
				res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
			}
		}
	}
	result := [][]int{}
	for triplet := range res {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}
	return result
}

func search(nums []int, l, r, target int) (int, bool) {
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m, true
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1, false
} */

// Brute force O(n^3 * 3log3)
/*func threeSum(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					res = addTriplet(res, triplet)
				}
			}
		}
	}
	return res
}

func addTriplet(list [][]int, triplet []int) [][]int {

	sort.Ints(triplet)

	for _, elem := range list {
		if reflect.DeepEqual(triplet, elem) {
			return list
		}
	}

	list = append(list, triplet)
	return list
}
*/
