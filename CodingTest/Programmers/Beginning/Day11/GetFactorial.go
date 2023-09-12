package main

func main() {
	println(solution(3628800))
	println(solution(7))
}

func solution(n int) int {
	var res int
	for i := 1; true; i++ {
		if n/i == 0 {
			res = i - 1
			break
		} else {
			n /= i
		}
	}
	return res
}
