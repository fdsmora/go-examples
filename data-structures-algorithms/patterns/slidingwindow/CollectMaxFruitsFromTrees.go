package slidingwindow

// leetcode problem 904
func totalFruit(fruits []int) int {
	const basketSize = 2
	if len(fruits) <= basketSize {
		return len(fruits)
	}
	var sum, max, s, e int
	basket := make(map[int]int, basketSize)

	for e = 0; e < len(fruits); e++ {
		_, ok := basket[fruits[e]]
		if len(basket) == 2 && !ok {
			sum = e - s
			if sum > max {
				max = sum
			}
			basket = make(map[int]int, basketSize)
			basket[fruits[e]] = 0
			basket[fruits[e-1]] = 0
			for s = e - 1; s >= 1 && fruits[s] == fruits[s-1]; s-- {
			}
		} else {
			basket[fruits[e]] = 0
		}
	}
	sum = e - s
	if sum > max {
		max = sum
	}
	return max
}

// 0  1  2  3  4  5
// 1  2  3  2  2
//                e
/*    s

B 1 2 | 3 2
sum  2 4
max  2
*/
