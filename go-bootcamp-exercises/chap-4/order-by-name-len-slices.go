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
	result := make([][]string, max)
	
	// 3. Fill result slice
	for _, n := range names {
		result[len(n)-1] = append(result[len(n)-1], n)
	}
		
	// 4. Print result
	fmt.Print(result)
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
