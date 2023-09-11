package main

import "fmt"

func main() {
	fmt.Println(good_solution([]int{1, 2, 3}, "right"))
	fmt.Println(solution([]int{1, 2, 3}, "right"))
	fmt.Println(good_solution([]int{4, 455, 6, 4, -1, 45, 6}, "left"))
	fmt.Println(solution([]int{4, 455, 6, 4, -1, 45, 6}, "left"))
}

func good_solution(numbers []int, direction string) []int {
	result := []int{}
	if direction == "left" {
		result = append(result, numbers[1:]...)
		result = append(result, numbers[0])
	} else {
		result = append(result, numbers[len(numbers)-1])
		result = append(result, numbers[:len(numbers)-1]...)
	}
	return result
}

func solution(numbers []int, direction string) []int {
	result := []int{}
	if direction == "left" {
		for i := 1; i < len(numbers); i++ {
			result = append(result, numbers[i])
		}
		result = append(result, numbers[0])
	} else {
		result = append(result, numbers[len(numbers)-1])
		for i := 0; i < len(numbers)-1; i++ {
			result = append(result, numbers[i])
		}
	}
	return result
}
