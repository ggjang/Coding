package main

import "fmt"

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
}

func solution(num_list []int, n int) [][]int {
	result := [][]int{{}}
	for i := 0; i < len(num_list); i += n {
		result = append(result, num_list[i:i+n])
	}
	result = append(result[1:])
	return result
}
