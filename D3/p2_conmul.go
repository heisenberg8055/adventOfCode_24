package D3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func ReadFile2() int64 {
	var ans int64 = 173731097
	file, err := os.OpenFile("./D3/input3.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	reff := ""
	for scanner.Scan() {
		reff += scanner.Text()
	}
	ans -= RegexMatchCon(reff)
	return ans
}

func RegexMatchCon(s string) int64 {
	var ans int64 = 0
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	arr := re.FindAllStringIndex(s, -1)
	fmt.Println(arr)
	for _, val := range arr {
		temp := ""
		for i := val[0] - 1; i <= val[1]+1; i++ {
			temp += string(s[i])
		}
		ans += RegexMatch(temp)
	}
	doregexp := regexp.MustCompile(`do\(\)`)
	do := doregexp.FindAllStringIndex(s, -1)
	dontregexp := regexp.MustCompile(`don't\(\)`)
	dont := dontregexp.FindAllStringIndex(s, -1)
	dol := do[(len(do) - 1)]
	last := dol[0]
	fmt.Println(dont, "test", do)
	for _, val := range dont {
		if last < val[0] {
			fmt.Println(val[0])
			temp := s[val[1]:]
			ans += RegexMatch(temp)
			break
		}
	}
	return ans
}

func Mullcon() int64 {
	return ReadFile2()
}
