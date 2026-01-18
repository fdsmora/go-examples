package slidingwindow

import "math"

func SmallestSubstringAllChars(chars []string, str string) string {
	validChars := [256]bool{}
	for _, c := range chars {
		if len(c) > 0 {
			validChars[c[0]] = true
		}
	}
	var s, e int
	minLen := math.MaxInt
	var minS, minE int
	for ; e < len(str); e++ {
		if !validChars[str[e]] {
			if currLen := e - s; currLen > 0 && currLen < minLen {
				minLen = currLen
				minS, minE = s, e
			}
			s = e + 1
		}
	}

	if currLen := e - s; currLen > 0 && currLen < minLen {
		minLen = currLen
		minS, minE = s, e
	}
	return str[minS:minE]
}

/*
{a b c d}
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
x y z m i o t u v
                  e
                  s

mS=8
mE=10
minL=2

{a b c d}
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
x a d c b a m i a b r   u v  d   d  d
                                      e
                             s
mS=8
mE=10
minL=2

*/
