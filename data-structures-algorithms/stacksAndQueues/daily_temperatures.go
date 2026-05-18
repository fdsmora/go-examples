package stacksandqueues

// leetcode problem 739
// My solution, time complexity is O(n) where n is the length of the input array, as we traverse the array once. Space complexity is O(n) in the worst case when the stack contains all indices of the input array.
func dailyTemperatures(temps []int) []int {
	if len(temps) == 1 {
		return []int{0}
	}
	st := []int{}
	res := make([]int, len(temps))

	for i := len(temps) - 1; i >= 0; i-- {
		c := temps[i]
		dist := 0
		for len(st) > 0 && c >= temps[st[len(st)-1]] {
			outIx := st[len(st)-1]
			st = st[:len(st)-1]
			dist += res[outIx]
		}
		if len(st) == 0 {
			dist = 0
		} else {
			dist += 1
		}
		res[i] = dist
		st = append(st, i)
	}
	return res
}
