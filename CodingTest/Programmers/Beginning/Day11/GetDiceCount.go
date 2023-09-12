package main

func main() {
	println(solution([]int{1, 1, 1}, 1))
	println(solution([]int{10, 8, 6}, 3))
}

func solution(box []int, n int) int {
	res := 1
	for _, v := range box {
		res *= v / n
	}
	return res
}
