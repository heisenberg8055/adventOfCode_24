package D1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ReadFile() ([]int, []int) {
	var arr1 = []int{}
	var arr2 = []int{}
	file, err := os.Open("./D1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempLine = scanner.Text()
		var (
			var1 string
			var2 string
		)
		var reff = true
		for _, arr := range tempLine {
			if arr == ' ' {
				reff = false
				continue
			}
			if reff {
				var1 += string(arr)
			} else {
				var2 += string(arr)
			}
		}
		val1, _ := strconv.Atoi(var1)
		val2, _ := strconv.Atoi(var2)
		arr1 = append(arr1, val1)
		arr2 = append(arr2, val2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr1, arr2
}

func DistanceDiff() int {
	var ans int
	arr1, arr2 := ReadFile()
	slices.Sort(arr1)
	slices.Sort(arr2)
	for i := 0; i < len(arr1); i++ {
		ans += absDiffInt(arr1[i], arr2[i])
	}
	return ans
}
