package main

import "fmt"

func main() {
	fmt.Println(solution(40))
	fmt.Println(solution(23))
}

func solution(age int) int {
	return 2023 - age
}
