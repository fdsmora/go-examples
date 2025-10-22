package miscelaneous

import (
	"math"
)

// Leetcode 5.
// O(n*k) where k is the length of the largest palindrome, but k is an upper bound.
// Not linear but better than O(n^2)
/* func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var maxLeft, maxRight int
	offset := 1
	i := 0
	count := 1
	for i < n {
		left, right := i-1, i+offset
		for left >= 0 && right < n {
			// expand pointers to discover palindrome
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		// the palindrome pattern breaks
		// set back to the palindrome limits
		left++
		right--
		// check and update max palindrome length
		if right-left+1 > maxRight-maxLeft+1 {
			maxLeft = left
			maxRight = right
		}
		if i < n-1 && s[i] == s[i+1] && count > 0 {
			offset = 2
			count--
		} else {
			offset = 1
			i++
			count = 1
		}
	}
	return s[maxLeft : maxRight+1]
} */

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	var maxLeft, maxRight int
	for i := 0; i < n; i++ {
		even := getPalindromeLength(s, i, i)
		odd := getPalindromeLength(s, i, i+1)
		maxLen := int(math.Max(float64(even), float64(odd)))

		if maxLen > maxRight-maxLeft+1 {
			maxLeft = i - (maxLen-1)/2
			maxRight = i + maxLen/2
		}
	}
	return s[maxLeft : maxRight+1]
}

func getPalindromeLength(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

/*
O(n^2) brute force solution
func longestPalindrome(s string) string {
    n := len(s)
    if n<=1{
        return s
    }

    var maxI, maxJ int

    for i:=0; i<n; i++{
        for j:=n-1; j>i; j--{
            if s[i]==s[j]{
                cI, cJ := checkPalindrome(s, i, j)
                if cJ-cI+1 > maxJ-maxI+1 {
                    maxI = cI
                    maxJ = cJ
                }
            }
        }
    }
    return s[maxI:maxJ+1]
}

func checkPalindrome(s string, i, j int) (int, int) {
    front, back := i, j
    for front < back {
        if s[front]!=s[back] {
            return 0,0
        }
        front++
        back--
    }
    return i, j
}
*/
