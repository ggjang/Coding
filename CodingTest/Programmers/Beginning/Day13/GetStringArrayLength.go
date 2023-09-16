package main

import "fmt"

func main() {
	fmt.Println(solution([]string{"We", "are", "the", "world!"}))
	fmt.Println(solution([]string{"I", "Love", "Programmers."}))
}

func solution(strlist []string) []int {
	result := []int{}
	for _, v := range strlist {
		result = append(result, len(v))
	}
	return result
}
