package main

import "fmt"

func main() {
	list1 := []int{1, 2, 3, 4, 5}
	result := solution(list1)
	fmt.Println(result)

	list2 := []int{1, 2, 100, -99, 1, 2, 3}
	result = solution(list2)
	fmt.Println(result)
}

func solution(numbers []int) []int {

	if result, isok := isValid(numbers); isok {
		return result
	} else {
		return nil
	}
}

func isValid(numbers []int) ([]int, bool) {
	if len(numbers) >= 1 && len(numbers) <= 1000 {
		for k, v := range numbers {
			if v < -10000 && v > 10000 {
				return nil, false
			} else {
				numbers[k] = v * 2
			}
		}
		return numbers, true
	} else {
		return nil, false
	}
}
