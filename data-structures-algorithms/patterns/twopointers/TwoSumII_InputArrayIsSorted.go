package twopointers

// leetcode problem 167
// runtime: O(n)
// space: O(1)
func TwoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	front, back := 0, len(numbers)-1
	for front = 0; front < len(numbers); front++ {
		diff := target - numbers[front]
		for diff < numbers[back] {
			back--
		}
		if numbers[front]+numbers[back] == target {
			break
		}
	}
	return []int{front + 1, back + 1}
}

func TwoSumSimpler(numbers []int, target int) []int {
	front, back := 0, len(numbers)-1

	for front < back {
		sum := numbers[front] + numbers[back]

		if sum == target {
			break
		}
		if sum > target {
			back--
		} else {
			front++
		}
	}
	return []int{front + 1, back + 1}
}
