package D16

import (
	"bufio"
	"os"
	"strings"
)

type cord struct {
	x     int
	y     int
	score int
	dir   rune
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
	ans := MaxInt
	arr := ReadFile()
	m := len(arr)
	n := len(arr[0])
	vis := make([][]int, m)
	for i := range vis {
		temp := make([]int, n)
		vis[i] = temp
	}
	vis[m-2][n-2] = 0
	queue := make([]cord, 0)
	queue = append(queue, cord{
		x:     m - 2,
		y:     n - 2,
		score: 0,
		dir:   69,
	})
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
	}
	return vis[1][n-2]
}
