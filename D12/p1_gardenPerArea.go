package D12

import (
	"bufio"
	"fmt"
	"os"
)

type cord struct {
	x int
	y int
}

func ReadFile() [][]rune {
	var arr [][]rune
	file, err := os.OpenFile("./D12/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := scanner.Text()
		var reff []rune
		for _, c := range temp {
			reff = append(reff, c)
		}
		arr = append(arr, reff)
	}
	return arr
}

func Help1() int {
	var ans int = 0
	arr := ReadFile()
	m := len(arr)
	n := len(arr[0])

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] {
				fmt.Print(i, j, " ")
				area, per := Bfs(i, j, m, n, arr[i][j], &vis, &arr)
				ans += area * per
			}
		}
	}
	return ans
}

func Bfs(i, j, m, n int, r rune, vis *[][]bool, arr *[][]rune) (int, int) {
	var (
		area   int = 0
		island int = 0
		neigh  int = 0
	)
	queue := make([]cord, 0)
	queue = append(queue, cord{x: i, y: j})
	for len(queue) != 0 {
		temp := queue[0]
		dx := temp.x
		dy := temp.y
		queue = queue[1:]
		if !(*vis)[dx][dy] {
			(*vis)[dx][dy] = true
			area++
			island++
			if dy+1 < n && (*arr)[dx][dy+1] == r {
				queue = append(queue, cord{dx, dy + 1})
				neigh++
			}
			if dx+1 < m && (*arr)[dx+1][dy] == r {
				queue = append(queue, cord{dx + 1, dy})
				neigh++
			}
			if dx > 0 && (*arr)[dx-1][dy] == r {
				queue = append(queue, cord{dx - 1, dy})
			}
			if dy > 0 && (*arr)[dx][dy-1] == r {
				queue = append(queue, cord{dx, dy - 1})
			}
		}
	}
	return area, island*4 - neigh*2
}
