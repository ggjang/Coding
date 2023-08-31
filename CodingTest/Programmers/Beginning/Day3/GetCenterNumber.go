package main

import "fmt"

func main() {
	array := []int{1, 2, 7, 10, 11}
	result := solution(array)
	fmt.Println(result)
	array = []int{9, -1, 0}
	result = solution(array)
	fmt.Println(result)
}

func solution(array []int) int {
	if isValid(array) {
		for i := 0; i < len(array); i++ {
			for ii := i + 1; ii < len(array); ii++ {
				if array[i] > array[ii] {
					temp := array[i]
					array[i] = array[ii]
					array[ii] = temp
				}
			}
			fmt.Println(array)
		}

		return array[len(array)/2]
	} else {
		return 0
	}
}

func isValid(array []int) bool {
	array_length := len(array)
	if array_length%2 == 0 {
		return false
	}
	if array_length < 0 || array_length > 100 {
		return false
	}

	for _, v := range array {
		if v <= -1000 || v >= 1000 {
			return false
		}
	}

	return true
}
