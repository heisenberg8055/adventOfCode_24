package D9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack []string

func (s *stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	val := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return val
}

func (s *stack) Push(v string) {
	(*s) = append((*s), v)
}

func (s *stack) Top() string {
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	} else {
		return ""
	}
}

func ReadFile() string {
	file, err := os.OpenFile("./D9/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return scanner.Text()
}

func srConst(c string, n int, st *stack, isDot *bool, cnt *int) string {
	var ans string = ""
	for n != 0 {
		(*st).Push(c)
		if *isDot {
			(*cnt)++
		}
		ans += c
		n--
	}
	return ans
}

func Help1() int {
	var st stack
	var isDot bool = false
	var ans int = 0
	reff := ""
	s := ReadFile()
	index := 0
	var cnt int = 0
	for i := range s {
		val, _ := strconv.Atoi(string(s[i]))
		if !isDot {
			ind := fmt.Sprint(index)
			index++
			reff += srConst(ind, val, &st, &isDot, &cnt)
		} else {
			reff += srConst(".", val, &st, &isDot, &cnt)
		}
		isDot = !isDot
	}
	c := cnt
	for i := 0; i < len(reff)-cnt; i++ {
		val := reff[i]
		if reff[i] == '.' {
			for st.Top() == "." {
				c--
				st.Pop()
			}
			var vall string = st.Pop()
			add, _ := strconv.Atoi(string(vall))
			ans += i * add
			c--
		} else {
			v, _ := strconv.Atoi(string(val))
			ans += i * v
		}
	}
	return ans
}
