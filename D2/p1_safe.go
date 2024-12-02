package d2

import (
	"log"
	"os"
)

func ReadFile() {
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
	scanner := 
}

func checkSafe(arr []int) bool {
	var (
		inc bool = true
		dec bool = true
	)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			inc = false
			break
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			dec = false
			break
		}
	}
	return inc && dec
}
