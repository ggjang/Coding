package main

func main() {
	println(solution(1234))
	println(solution(930211))
}

func solution(n int) int {
	result := 0
	for n > 0 {
		result += n % 10
		n /= 10
	}
	return result
}
