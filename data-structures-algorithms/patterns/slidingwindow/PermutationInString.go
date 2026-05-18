package slidingwindow

// leetcode problem 567
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	var counter []int
	var sum int
	counter, sum = getCounter(s1)

	front := 0
	origSum := sum

	for front+len(s1) <= len(s2) {
		for i := front; i < front+len(s1); i++ {
			c := s2[i]
			if counter[c-'a'] == 0 {
				break
			}
			counter[c-'a']--
			sum--
		}
		if sum == 0 {
			return true
		}
		if sum != origSum {
			counter, sum = getCounter(s1)
		}
		front++
	}
	return false
}

func getCounter(s1 string) ([]int, int) {
	var (
		sum     int
		counter = make([]int, 26)
	)
	for _, c := range s1 {
		counter[c-'a'] += 1
		sum += 1
	}
	return counter, sum
}
