package main

func main() {
	println(solution([]int{7, 77, 17}))
	println(solution([]int{10, 29}))
}

func solution(array []int) int {
	result := 0
	for _, v := range array {
		for v > 0 {
			if digit := v % 10; digit%7 == 0 && digit != 0 {
				result++
			}
			v /= 10
		}
	}
	return result
}
