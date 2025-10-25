package arrayshash

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	anagrams := make(map[[26]int][]string)
	for _, s := range strs {
		frequency := getFrequency(s)
		if anagrams[frequency] == nil {
			anagrams[frequency] = []string{}
		}
		anagrams[frequency] = append(anagrams[frequency], s)
	}
	for _, list := range anagrams {
		result = append(result, list)
	}
	return result
}

func getFrequency(s1 string) [26]int {
	charCount := [26]int{}
	for _, c := range s1 {
		charCount[c-'a']++
	}
	return charCount
}
