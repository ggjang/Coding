package main

import (
	"strconv"
	"strings"
)

func main() {
	println(solution("1 2 Z 3"))
	println(solution("10 20 30 40"))
	println(solution("10 Z 20 Z 1"))
	println(solution("10 Z 20 Z"))
	println(solution("-1 -2 -3 Z"))
}

func solution(s string) int {
	res := 0
	toks := strings.Split(s, " ")
	for k, v := range toks {
		if v == "Z" {
			value, _ := strconv.Atoi(toks[k-1])
			res -= value
		} else {
			value, _ := strconv.Atoi(v)
			res += value
		}
	}
	return res
}
