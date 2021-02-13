package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	// 1. Get max name length
	max := getMaxLen(names)
	
	// 2. Allocate result slice
	result := make([][]string, max+1)
	
	// 3. Fill result slice
	for _, n := range names {
		appendVal(result, n)
	}
		
	// 4. Print result
	fmt.Print(result)
}

func appendVal(arr [][]string, str string) {
	slen := len(str)
	arr[slen] = append(arr[slen], str) 
}

func getMaxLen(arr []string) int {
	var max int
	for _, n := range names {
		if curlen := len(n); curlen > max {
			max = curlen
		}
	}
	return max
}

