package main

func main() {
	println(solution([]int{1, 1, 2, 3, 4, 5}, 1))
	println(solution([]int{0, 2, 3, 4}, 1))
}

func solution(array []int, n int) int {
	result := 0
	for _, v := range array {
		if v == n {
			result++
		}
	}

	return result
}
