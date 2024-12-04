package D4

import "fmt"

func Help2() int {
	ans := 0
	arr := ReadFile()
	var m int = len(arr)
	var n int = len(arr[0])
	q := make([]cord, 0)
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if arr[i][j] == 65 {
				q = append(q, cord{x: i, y: j})
			}
		}
	}
	for _, temp := range q {
		x := temp.x
		y := temp.y
		var cross1 string = string(arr[x-1][y+1]) + string(arr[x][y]) + string(arr[x+1][y-1])
		var cross2 string = string(arr[x-1][y-1]) + string(arr[x][y]) + string(arr[x+1][y+1])
		fmt.Println(cross1, cross2)
		if (cross1 == "MAS" || cross1 == "SAM") && (cross2 == "MAS" || cross2 == "SAM") {
			ans++
		}
	}
	return ans
}
