package D3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadFile() int64 {
	var ans int64
	file, err := os.OpenFile("./D3/input3.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans += RegexMatch(scanner.Text())
	}
	return ans
}

func RegexMatch(s string) int64 {
	var ans int64 = 0
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	arr := re.FindAllString(s, -1)
	for _, val := range arr {
		re1 := regexp.MustCompile(`[0-9]+`)
		brr := re1.FindAllString(val, -1)
		val1, _ := strconv.Atoi(brr[0])
		val2, _ := strconv.Atoi(brr[1])
		ans += int64(val1) * int64(val2)
	}
	return ans
}

func Mull() int64 {
	return ReadFile()
}
