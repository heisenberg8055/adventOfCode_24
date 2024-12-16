package D12

func Help2() int {
	arr := ReadFile()
	var ans int = 0
	m := len(arr)
	n := len(arr[0])

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] {

			}
		}
	}
	return ans
}
