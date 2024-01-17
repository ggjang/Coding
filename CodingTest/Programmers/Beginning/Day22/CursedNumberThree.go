package main

func main() {
	println(solution(40))
}

func solution(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result++
		for result%100-result%10 == 30 || result%3 == 0 || result%10 == 3 {
			result++
		}
	}
	return result
}
