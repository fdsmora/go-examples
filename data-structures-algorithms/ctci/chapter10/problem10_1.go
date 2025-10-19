package chapter10

// CTCI problem 10.1
func sortArrays(A, B []int, countA, countB int) {
	lastIdx := countA + countB - 1
	b := countB - 1
	a := countA - 1

	for b >= 0 {
		if a >= 0 && A[a] > B[b] {
			A[lastIdx] = A[a]
			a--
		} else {
			A[lastIdx] = B[b]
			b--
		}
		lastIdx--
	}
}

//  0 1 2 3 4  5   6  7  8  9  10 11
//A 3 3 4 5 7 8  9  11  13  15  19
// a
//    i
//B 2 4 7 19
//  b
// lastIdx : 10
// a: 6
// b: 2
