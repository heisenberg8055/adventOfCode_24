package D5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([]string, []string) {
	file, _ := os.OpenFile("./D5/input.txt", os.O_RDONLY, 0644)
	scanner := bufio.NewScanner(file)
	var rules []string
	for scanner.Scan() {
		temp := scanner.Text()
		if len(temp) == 0 {
			break
		}
		rules = append(rules, temp)
	}
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return rules, arr
}

func chk(arr []string, brr string) bool {
	for _, a := range arr {
		if a == brr {
			return true
		}
	}
	return false
}

func Help() int {
	var ans int = 0
	r, arr := ReadFile()
	mp := make(map[string][]string, 0)
	for _, s := range r {
		s1 := string(s[0]) + string(s[1])
		s2 := string(s[3]) + string(s[4])
		mp[s2] = append(mp[s2], s1)
	}

	for _, a := range arr {
		curr := strings.Split(a, ",")
		n := len(curr)
		var add bool = true
		for i := 0; i < n; i++ {
			if !add {
				break
			}
			for j := i + 1; j < n; j++ {
				if chk(mp[curr[i]], curr[j]) {
					add = false
					break
				}
			}
		}
		if add {
			val, _ := strconv.Atoi(curr[n/2])
			ans += val
		}
	}

	return ans
}
