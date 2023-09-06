package main

func main() {
	println(solution(10))
	println(goodSolution(10))
}

func solution(n int) (sum int) {
	for i := 0; i <= n; i += 2 {
		sum += i
	}
	return sum
}

func goodSolution(n int) int {
	return (1 + n/2) * (n / 2)
}
