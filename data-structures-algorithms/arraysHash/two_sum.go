
// Hashmap 2 pass
/*func twoSum(nums []int, target int) []int {
    indices := make(map[int]int, len(nums))

    for i, v := range nums {
        indices[v]=i
    }

    i, j := 0,0
    var ok bool
    for ; i<len(nums); i++ {
        if j, ok = indices[target - nums[i]]; ok  && j != i {
            break
        }
    }
    return []int{i, j}
}*/

// Hashmap 1 pass
func twoSum(nums []int, target int) []int {
	indices := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, exists := indices[target-v]; exists {
			return []int{j, i}
		}

		indices[v] = i
	}
	return []int{}
}


