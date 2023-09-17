package main

import "fmt"

func main() {
	fmt.Println(solution(24))
	fmt.Println(solution(29))
}

func solution(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
