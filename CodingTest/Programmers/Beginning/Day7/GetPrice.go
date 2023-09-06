package main

func main() {
	println(solution(10, 3))
	println(solution(64, 6))
}

func solution(n int, k int) int {
	return n*12000 + (k-n/10)*2000
}
