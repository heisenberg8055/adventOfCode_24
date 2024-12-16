package D16

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cord struct {
	x   int
	y   int
	dir rune
}

func ReadFile() [][]string {
	var arr [][]string
	file, err := os.OpenFile("./D16/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		reff := strings.Split(temp, "")
		arr = append(arr, reff)
	}
	return arr
}

func Help1() int {
	const MaxInt = int(^uint(0) >> 1)
	arr := ReadFile()
	m := len(arr)
	n := len(arr[0])
	vis := make([][]int, m)
	for i := range vis {
		temp := make([]int, n)
		for i := range temp {
			temp[i] = 100000
		}
		vis[i] = temp
	}
	vis[m-2][1] = 0
	queue := make([]cord, 0)
	queue = append(queue, cord{
		x:   m - 2,
		y:   1,
		dir: 69,
	})
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		dx := temp.x
		dy := temp.y
		currDir := temp.dir
		if dx > 0 && arr[dx-1][dy] != "#" {
			var dist int
			switch currDir {
			case 'E':
				dist = 1001
			case 'S':
				dist = 2001
			case 'N':
				dist = 1
			case 'W':
				dist = 1001
			}
			if vis[dx-1][dy] > vis[dx][dy]+dist {
				vis[dx-1][dy] = vis[dx][dy] + dist
				queue = append(queue, cord{x: dx - 1, y: dy, dir: 'N'})
			}
		}
		if dy > 0 && arr[dx][dy-1] != "#" {
			var dist int
			switch currDir {
			case 'E':
				dist = 2001
			case 'S':
				dist = 1001
			case 'N':
				dist = 1001
			case 'W':
				dist = 1
			}
			if vis[dx][dy-1] > vis[dx][dy]+dist {
				vis[dx][dy-1] = vis[dx][dy] + dist
				queue = append(queue, cord{x: dx, y: dy - 1, dir: 'W'})
			}
		}
		if dy < n-1 && arr[dx][dy+1] != "#" {
			var dist int
			switch currDir {
			case 'E':
				dist = 1
			case 'S':
				dist = 1001
			case 'N':
				dist = 1001
			case 'W':
				dist = 2001
			}
			if vis[dx][dy+1] > vis[dx][dy]+dist {
				vis[dx][dy+1] = vis[dx][dy] + dist
				queue = append(queue, cord{x: dx, y: dy + 1, dir: 'E'})
			}
		}
		if dx < m-1 && arr[dx+1][dy] != "#" {
			var dist int
			switch currDir {
			case 'E':
				dist = 1001
			case 'S':
				dist = 1
			case 'N':
				dist = 2001
			case 'W':
				dist = 1001
			}
			if vis[dx+1][dy] > vis[dx][dy]+dist {
				vis[dx+1][dy] = vis[dx][dy] + dist
				queue = append(queue, cord{x: dx + 1, y: dy, dir: 'W'})
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(vis[i][j], "\t")
		}
		fmt.Println()
	}
	return vis[1][n-2]
}
