package slidingwindow

import "math"

// leetcode problem 904
// runtime O(2n) -> O(n), as e visits n fruits, and in the worst case, s goes back from e-1 to 1,
// so when e = n-1, s visits almost all n elements backwards, adding another n so O(n + n)
// space is O(1) as the basket only holds 2 items at most
func totalFruit(fruits []int) int {
	const basketSize = 2
	if len(fruits) <= basketSize {
		return len(fruits)
	}
	var sum, max, s, e int
	basket := make(map[int]bool, basketSize)

	for e = 0; e < len(fruits); e++ {
		if len(basket) == 2 && !basket[fruits[e]] {
			sum = e - s
			max = int(math.Max(float64(max), float64(sum)))
			basket = make(map[int]bool, basketSize)
			basket[fruits[e]] = true
			basket[fruits[e-1]] = true
			for s = e - 1; s >= 1 && fruits[s] == fruits[s-1]; s-- {
			}
		} else {
			basket[fruits[e]] = true
		}
	}
	sum = e - s

	return int(math.Max(float64(max), float64(sum)))
}

// 0  1  2  3  4  5
// 1  2  3  2  2
//                e
/*    s

B 1 2 | 3 2
sum  2 4
max  2
*/
