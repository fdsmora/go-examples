package arrayshash

//import "sort"

// LC 347
// O(n) runtime, O(n) space, using frequency buckets
func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	buckets := make([][]int, len(nums)+1)

	for _, v := range nums {
		freq[v]++
	}

	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}

	res := []int{}
	for i := len(buckets) - 1; i > 0; i-- {
		if len(buckets[i]) > 0 {
			for _, v := range buckets[i] {
				if k > 0 {
					res = append(res, v)
					k--
				} else {
					return res
				}
			}
		}
	}
	return res
}

// O(nlogn) solution O(n) space
/*
func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	ordered := [][]int{}
	for k, v := range count {
		ordered = append(ordered, []int{k, v})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i][1] > ordered[j][1]
	})
	res := []int{}
	for _, v := range ordered[0:k] {
		res = append(res, v[0])
	}
	return res
}
*/
