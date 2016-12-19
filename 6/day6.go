package day6

import (
	"bytes"
	"fmt"
	"strings"
)

func Run(input string, task int) {

	lines := strings.Split(input, "\n")

	occurency := make([]map[rune]int, len(lines[0]))
	for i := 0; i < len(occurency); i++ {
		occurency[i] = map[rune]int{}
	}

	for _, line := range lines {
		for i, r := range line {
			if val, ok := occurency[i][r]; ok {
				occurency[i][r] = val + 1
			} else {

				occurency[i][r] = 1
			}

		}
	}

	compareFunc := func(i, j int) bool { return i > j }
	initialValue := -1
	if task == 2 {
		compareFunc = func(i, j int) bool { return i < j }
		initialValue = 999
	}

	var buffer bytes.Buffer
	for _, occ := range occurency {
		r, i := extremeEntryInMap(occ, initialValue, compareFunc)
		if i == -1 {
			panic("aaa")
		} else {
			buffer.WriteRune(r)
		}
	}
	fmt.Printf("password: %s\n", buffer.String())
}

func extremeEntryInMap(m map[rune]int, initialValue int, compare func(i, j int) bool) (r rune, i int) {
	r = 'a'
	i = initialValue
	for k, v := range m {
		if compare(v, i) {
			r, i = k, v
		}
	}
	return
}
