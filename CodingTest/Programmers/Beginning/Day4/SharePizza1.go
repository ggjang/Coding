package main

func main() {
	var result int
	result = solution(7)
	println(result)

	result = solution(1)
	println(result)

	result = solution(15)
	println(result)
}

func solution(n int) int {
	if isValid(n) {
		share := n / 7
		if n%7 != 0 {
			return share + 1
		} else {
			return share
		}

	}
	return 0
}

func isValid(n int) bool {
	if n >= 1 && n <= 100 {
		return true
	}
	return false
}
