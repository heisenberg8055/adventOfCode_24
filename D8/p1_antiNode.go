package D8

import (
	"bufio"
	"os"
	"strings"
)

type cord struct {
	x int
	y int
}

func ReadFile() [][]string {
	file, err := os.OpenFile("./D8/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	var arr [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		reff := strings.Split(temp, "")
		arr = append(arr, reff)
	}
	return arr
}

func Help1() int {
	var ans int = 0
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
				reff1 := calc(cord1, cord2)
				reff2 := calc(cord2, cord1)
				vis[reff1] = true
				vis[reff2] = true
			}
		}
	}
	for key := range vis {
		dx := key.x
		dy := key.y
		if dx >= 0 && dx < m && dy >= 0 && dy < n {
			ans++
		}
	}
	return ans
}

func calc(cord1 cord, cord2 cord) cord {
	var ans cord
	ans.x = 2*cord1.x - cord2.x
	ans.y = 2*cord1.y - cord2.y
	return ans
}
