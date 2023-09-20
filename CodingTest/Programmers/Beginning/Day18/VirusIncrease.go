package main

func main() {
	println(solution(2, 10))
	println(solution(7, 15))
}

func solution(n int, t int) int {
	result := n
	for i := 0; i < t; i++ {
		result = result * 2
	}
	return result
}
