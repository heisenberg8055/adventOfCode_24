package D2

func Calc2() int {
	var ans int = 0
	arr := ReadFile()
	for _, a := range arr {
		if checkSafeOne(a) {
			ans++
		}
	}
	return ans
}

func checkSafeOne(arr []int) bool {
	if checkSafe(arr) {
		return true
	} else {
		n := len(arr)
		for i := 0; i < n; i++ {
			brr := []int{}
			for j:= 0; j < n; j++ {
				if (i != j) {
					brr = append(brr, arr[j])
				}
			}
			if checkSafe(brr) {
				return true
			}
		}
	}
	return false
}
