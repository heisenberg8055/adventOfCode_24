package D12

import (
	"bufio"
	"os"
)

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

	return ans
}
