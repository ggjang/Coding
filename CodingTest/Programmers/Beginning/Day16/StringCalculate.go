package main

import (
	"strconv"
	"strings"
)

func main() {
	println(solution("3 + 4"))
}

func solution(my_string string) int {
	toks := strings.Split(my_string, " ")
	res := 0
	// true(+), false(-)
	op := true
	for _, v := range toks {
		value, err := strconv.Atoi(v)
		if err != nil {
			if v == "+" {
				op = true
			} else {
				op = false
			}
			continue
		}

		if op {
			res += value
		} else {
			res -= value
		}

	}
	return res
}
