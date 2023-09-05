package main

import "fmt"

func main() {
	fmt.Println(solution(5500))
	fmt.Println(solution(15000))
}

func solution(money int) []int {
	cupCount := money / 5500
	remain := money % 5500
	return []int{cupCount, remain}
}
