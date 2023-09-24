package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	println(solution("1"))
}

func solution(polynomial string) string {
	toks := strings.Split(polynomial, " ")
	x_num := 0
	num := 0
	for _, v := range toks {
		if strings.Contains(v, "x") {
			only_num := strings.Replace(v, "x", "", -1)
			if len(only_num) == 0 {
				x_num += 1
				continue
			}
			value, err := strconv.Atoi(only_num)
			if err == nil {
				x_num += value
			}
			continue
		}

		value, err := strconv.Atoi(v)
		if err == nil {
			num += value
		}
	}

	var result string
	if x_num > 1 {
		result += fmt.Sprintf("%dx", x_num)
	} else if x_num == 1 {
		result += "x"
	}

	if strings.Contains(polynomial, "x") {
		if num > 0 {
			result += fmt.Sprintf(" + %d", num)
		} else if x_num < 0 {
			result += fmt.Sprintf(" - %d", -num)
		}
	} else {
		if num > 0 {
			result += fmt.Sprintf("%d", num)
		} else if x_num < 0 {
			result += fmt.Sprintf("%d", -num)
		}
	}

	return result
}
