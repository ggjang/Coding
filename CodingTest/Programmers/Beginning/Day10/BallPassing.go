package main

func main() {
	println(solution([]int{1, 2, 3, 4}, 2))
	println(solution([]int{1, 2, 3, 4, 5, 6}, 5))
	println(solution([]int{1, 2, 3}, 3))
}

func solution(numbers []int, k int) int {
	passing_index := k * 2
	return numbers[(passing_index-2)%len(numbers)]
}
