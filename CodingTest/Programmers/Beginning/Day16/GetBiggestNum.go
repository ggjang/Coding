package main

import "fmt"

func main() {
	fmt.Println(solution([]int{1, 8, 3}))
	fmt.Println(solution([]int{9, 10, 11, 8}))
}

func solution(array []int) []int {
	temp := make(map[int]int)
	for k, v := range array {
		temp[v] = k
	}

	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	return []int{array[len(array)-1], temp[array[len(array)-1]]}
}
