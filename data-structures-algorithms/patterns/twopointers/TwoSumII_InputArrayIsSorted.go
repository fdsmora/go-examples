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

// Binary search solution
/*
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i, v := range numbers {
		want := target - v
		if want == v{
			continue
		}
		if j, found := bs(numbers, want, i+1, n-1); found {
			return []int{i+1,j+1}
		}
	}
	return nil
}

func bs(nums []int, want, s, e int) (int, bool){
	for s <= e {
		m := (s+e)/2
		if want == nums[m]{
			return m, true
		}
		if want < nums[m] {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return 0, false
}
*/
