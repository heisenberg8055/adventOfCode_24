package D16

import (
	"bufio"
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
			temp[i] = MaxInt
		}
		vis[i] = temp
	}
	vis[m-2][1] = 0
	q := New[cord, int64](MaxHeap)
	q.Put(cord{x: m - 2, y: 1, dir: 69}, int64(0))
	for !q.IsEmpty() {
		temp := q.Get()
		dx := temp.Value.x
		dy := temp.Value.y
		currDir := temp.Value.dir
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
				q.Put(cord{x: dx - 1, y: dy, dir: 'N'}, int64(vis[dx-1][dy]))
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
				q.Put(cord{x: dx, y: dy - 1, dir: 'W'}, int64(vis[dx][dy-1]))
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
				q.Put(cord{x: dx, y: dy + 1, dir: 'E'}, int64(vis[dx][dy+1]))
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
				q.Put(cord{x: dx + 1, y: dy, dir: 'S'}, int64(vis[dx+1][dy]))
			}
		}
	}
	return vis[1][n-2]
}
