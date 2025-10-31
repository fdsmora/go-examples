package slidingwindow

// LC 424
// My solution
func characterReplacement(s string, k int) int {
	var left, maxLen int
	var chCount = map[byte]int{}

	for right := 0; right < len(s); right++ {
		chCount[s[right]]++
		winLen := right - left + 1
		mostRepCount := getMostRepeatedCount(chCount)
		totalReps := winLen - mostRepCount
		if totalReps <= k {
			maxLen = max(maxLen, winLen)
		} else {
			for totalReps > k {
				chCount[s[left]]--
				left++
				winLen--
				mostRepCount := getMostRepeatedCount(chCount)
				totalReps = winLen - mostRepCount
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMostRepeatedCount(chCount map[byte]int) int {
	maxCount := 0
	for _, v := range chCount {
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount
}
