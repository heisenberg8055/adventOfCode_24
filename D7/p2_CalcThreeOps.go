package D7

import "strconv"

func Help2() int {
	var ans int = 0
	arr, brr := ReadFile()
	n := len(arr)
	for i := 0; i < n; i++ {
		if recur2(1, arr[i], brr[i], brr[i][0]) {
			ans += arr[i]
		}
	}
	return ans
}

func recur2(i int, val int, arr []int, calc int) bool {
	if i == len(arr) {
		return calc == val
	}
	p := recur2(i+1, val, arr, calc+arr[i])
	b := recur2(i+1, val, arr, calc*arr[i])
	c := recur2(i+1, val, arr, combi(calc, arr[i]))
	return p || b || c
}

func combi(a int, b int) int {
	aS := strconv.Itoa(a)
	bS := strconv.Itoa(b)
	ans, _ := strconv.Atoi(aS + bS)
	return ans
}
