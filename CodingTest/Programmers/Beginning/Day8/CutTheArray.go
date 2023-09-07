package main

import "fmt"

func main() {
	result := solution([]int{1, 2, 3, 4, 5}, 1, 3)
	fmt.Println(result)
	result = solution([]int{1, 3, 5}, 1, 2)
	fmt.Println(result)
}

func solution(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}
