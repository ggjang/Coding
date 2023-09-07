package main

func main() {
	println(solution(20))
	println(solution(100))
}

func solution(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}

	return result
}
