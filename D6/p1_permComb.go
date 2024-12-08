package D6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() {
	f, err := os.OpenFile("./D6/input.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	mp := make(map[int64][]int64)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ":")
		key := arr[0]
		keyConv, _ := strconv.Atoi(key)
		k := int64(keyConv)
		value := arr[1]
		var a []int64
		val := strings.Split(value, " ")
		for i := range val {
			if i == 0 {
				continue
			}
			conv, _ := strconv.Atoi(val[i])
			a = append(a, int64(conv))
		}
		mp[k] = a
	}
	fmt.Print(mp)
}
