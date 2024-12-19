package D8

func Help2() int {
	arr := ReadFile()
	mp := make(map[string][]cord, 0)
	m := len(arr)
	n := len(arr[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != "." {
				mp[arr[i][j]] = append(mp[arr[i][j]], cord{x: i, y: j})
			}
		}
	}
	vis := make(map[cord]bool)
	for _, val := range mp {
		nums := val
		size := len(nums)
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				cord1, cord2 := nums[i], nums[j]
				calc2(cord1, cord2, m, n, &vis)
			}
		}
	}

	return len(vis)
}

func isValid(x, y, m, n int) bool {
	if x >= 0 && y >= 0 && x < m && y < n {
		return true
	} else {
		return false
	}
}

func calc2(cord1 cord, cord2 cord, m int, n int, vis *map[cord]bool) {
	dx1, dx2, dy1, dy2 := cord1.x, cord2.x, cord1.y, cord2.y
	for isValid(dx1, dy1, m, n) {
		(*vis)[cord{x: dx1, y: dy1}] = true
		temp1, temp2 := dx1, dy1
		dx1 = dx1*2 - dx2
		dy1 = dy1*2 - dy2
		dx2, dy2 = temp1, temp2
	}
	dx1, dx2, dy1, dy2 = cord1.x, cord2.x, cord1.y, cord2.y
	for isValid(dx2, dy2, m, n) {
		(*vis)[cord{x: dx2, y: dy2}] = true
		temp1, temp2 := dx2, dy2
		dx2 = dx2*2 - dx1
		dy2 = dy2*2 - dy1
		dx1, dy1 = temp1, temp2
	}
}
