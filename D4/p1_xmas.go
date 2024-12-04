package D4

import (
	"bufio"
	"os"
)

type cord struct {
	x int
	y int
}

func ReadFile() [][]rune {
	f, err := os.OpenFile("./D4/input4.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	arr := make([][]rune, 0)

	for scanner.Scan() {
		s := scanner.Text()
		brr := make([]rune, 0)
		for _, val := range s {
			brr = append(brr, val)
		}
		arr = append(arr, brr)
	}
	// for _, x := range arr {
	// 	for _, y := range x {
	// 		fmt.Print(y)
	// 	}
	// 	fmt.Println()
	// }
	return arr
}

func Help() int {
	ans := 0
	arr := ReadFile()
	var m int = len(arr)
	var n int = len(arr[0])
	q := make([]cord, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 88 {
				q = append(q, cord{x: i, y: j})
			}
		}
	}
	for _, temp := range q {
		x := temp.x
		y := temp.y
		if x+3 < m && arr[x][y] == 88 && arr[x+1][y] == 77 && arr[x+2][y] == 65 && arr[x+3][y] == 83 {
			ans++
		}
		if x-3 >= 0 && arr[x][y] == 88 && arr[x-1][y] == 77 && arr[x-2][y] == 65 && arr[x-3][y] == 83 {
			ans++
		}
		if y-3 >= 0 && arr[x][y] == 88 && arr[x][y-1] == 77 && arr[x][y-2] == 65 && arr[x][y-3] == 83 {
			ans++
		}
		if y+3 < n && arr[x][y] == 88 && arr[x][y+1] == 77 && arr[x][y+2] == 65 && arr[x][y+3] == 83 {
			ans++
		}
		if x+3 < m && y+3 < n && arr[x][y] == 88 && arr[x+1][y+1] == 77 && arr[x+2][y+2] == 65 && arr[x+3][y+3] == 83 {
			ans++
		}
		if x+3 < m && y-3 >= 0 && arr[x][y] == 88 && arr[x+1][y-1] == 77 && arr[x+2][y-2] == 65 && arr[x+3][y-3] == 83 {
			ans++
		}
		if x-3 >= 0 && y-3 >= 0 && arr[x][y] == 88 && arr[x-1][y-1] == 77 && arr[x-2][y-2] == 65 && arr[x-3][y-3] == 83 {
			ans++
		}
		if x-3 >= 0 && y+3 < n && arr[x][y] == 88 && arr[x-1][y+1] == 77 && arr[x-2][y+2] == 65 && arr[x-3][y+3] == 83 {
			ans++
		}
	}
	return ans
}
