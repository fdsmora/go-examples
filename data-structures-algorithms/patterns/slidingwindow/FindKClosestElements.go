package slidingwindow

// leetcode problem 658
// My solution, time complexity is O(n) where n is the length of the input array, as we traverse the array once. Space complexity is O(1) as we are using a constant amount of space to store the indices and the result.
func findClosestElements(arr []int, k int, x int) []int {
	s, e, f := 0, k-1, k

	for f < len(arr) {
		if abs(arr[f]-x) < abs(arr[s]-x) {
			e = f
			s = e - k + 1
		}
		f++
	}
	return arr[s : e+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
