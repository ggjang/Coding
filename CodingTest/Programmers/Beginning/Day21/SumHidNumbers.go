package main

import (
	"strconv"
)

func main() {
	println(solution("aAb1B2cC34oOp"))
	println(solution("1a2b3c4d123Z"))
}

func solution(my_string string) int {
	result := 0
	temp := []byte{}
	for i := 0; i < len(my_string); i++ {
		if my_string[i] >= '0' && my_string[i] <= '9' {
			temp = append(temp, my_string[i])
		} else {
			value, _ := strconv.Atoi(string(temp))
			result += value
			temp = []byte{}
		}
	}

	if len(temp) != 0 {
		value, _ := strconv.Atoi(string(temp))
		result += value
	}

	return result
}
