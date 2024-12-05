package D5

import "strconv"

func swap(slice []string, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func Chk(mp map[string]int, arr []string) bool {
	n := len(arr)
	var swp bool = false
	for i := 0; i < n && !swp; i++ {
		for j := i + 1; j < n && !swp; j++ {
			s1 := string(arr[j]) + "|" + string(arr[i])
			if mp[s1] != 0 {
				swap(arr, i, j)
				swp = true
			}
		}
	}
	return swp
}

func Help2() int {
	var ans int = 0
	r, _ := ReadFile()
	mp := make(map[string]int, 0)
	for _, s := range r {
		mp[s]++
	}

	arr, _ := Result()

	for _, a := range arr {
		for Chk(mp, a) {

		}
		n := len(a)
		reff, _ := strconv.Atoi(string(a[n/2]))
		ans += reff
	}

	return ans
}
