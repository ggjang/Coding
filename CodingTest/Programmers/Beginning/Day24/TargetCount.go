package main

import (
	"strconv"
)

func main() {
	println(solution(1, 13, 1))
	//println(solution(10, 50, 5))
	//println(solution(3, 10, 2))
}

func solution(i int, j int, k int) int {
	result := 0
	var target rune = rune(k + 48)
	for ; i <= j; i++ {
		str := strconv.Itoa(i)
		for _, v := range str {
			if v == target {
				result++
			}
		}
	}

	return result
}
