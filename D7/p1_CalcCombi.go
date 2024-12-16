package D7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([]int, [][]int) {
	file, err := os.OpenFile("./D7/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	var total []int
	var val [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		arr := strings.Split(temp, ":")
		reff, _ := strconv.Atoi(arr[0])
		total = append(total, reff)
		brr := strings.Split(arr[1], " ")
		brr = brr[1:]
		var values []int
		for i := range brr {
			val, _ := strconv.Atoi(brr[i])
			values = append(values, val)
		}
		val = append(val, values)
	}
	return total, val
}

func Help1() int {
	var ans int = 0
	arr, brr := ReadFile()
	n := len(arr)
	for i := 0; i < n; i++ {
		if recur(1, arr[i], brr[i], brr[i][0]) {
			ans += arr[i]
		}
	}
	return ans
}

func recur(i int, val int, arr []int, calc int) bool {
	if i == len(arr) {
		return calc == val
	}
	p := recur(i+1, val, arr, calc+arr[i])
	b := recur(i+1, val, arr, calc*arr[i])
	return p || b
}
