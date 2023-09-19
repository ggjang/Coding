package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution([]string{"3 - 4 = -3", "5 + 6 = 11"}))
}

func solution(quiz []string) []string {
	result := []string{}
	for _, v := range quiz {
		if isCorrect(v) {
			result = append(result, "O")
		} else {
			result = append(result, "X")
		}

	}
	return result
}

func isCorrect(prob string) bool {
	// true(+), false(-)
	toks := strings.Split(prob, " ")
	num1, _ := strconv.Atoi(toks[0])
	num2, _ := strconv.Atoi(toks[2])
	result, _ := strconv.Atoi(toks[4])
	if toks[1] == "+" {
		return result == (num1 + num2)
	}

	return result == (num1 - num2)
}
