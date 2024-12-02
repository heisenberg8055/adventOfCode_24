package D1

func SimilarityDiff() int {
	var ans int
	arr1, arr2 := ReadFile()
	var mp = make(map[int]int)
	for _, val := range arr2 {
		mp[val]++
	}
	for _, val := range arr1 {
		ans += mp[val] * val
	}
	return ans
}
