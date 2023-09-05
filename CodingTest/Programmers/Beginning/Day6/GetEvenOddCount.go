package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	result := solution(array)
	fmt.Println(result)

	array = []int{1, 3, 5, 7}
	result = solution(array)
	fmt.Println(result)
}

func solution(num_list []int) []int {
	even := 0
	odd := 0
	for _, v := range num_list {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return []int{even, odd}
}
