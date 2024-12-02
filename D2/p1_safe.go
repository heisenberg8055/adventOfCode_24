package D2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ReadFile() [][]int {
	file, err := os.OpenFile("./D2/input2.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	a := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Fields(s)
		irr := make([]int, len(arr))
		for i, val := range arr {
			str, _ := strconv.Atoi(val)
			irr[i] = str
		}
		a = append(a, irr)
	}
	return a
}

func Calc() int {
	var ans int = 0
	arr := ReadFile()
	for _, a := range arr {
		if checkSafe(a) {
			ans++
		}
	}
	return ans
}

func checkSafe(arr []int) bool {
	var (
		inc bool = true
		dec bool = true
	)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] || (absDiffInt(arr[i-1], arr[i]) < 1 || absDiffInt(arr[i-1], arr[i]) > 3) {
			inc = false
			break
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] || (absDiffInt(arr[i-1], arr[i]) < 1 || absDiffInt(arr[i-1], arr[i]) > 3) {
			dec = false
			break
		}
	}
	return inc || dec
}
