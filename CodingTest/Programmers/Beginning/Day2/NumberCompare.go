package main

func main() {
	result := solution(2, 3)
	println(result)

	result = solution(11, 11)
	println(result)

	result = solution(7, 99)
	println(result)
}

func solution(num1 int, num2 int) int {
	if isValid(num1) && isValid(num2) {
		if num1 == num2 {
			return 1
		} else {
			return -1
		}
	} else {
		return 0
	}
}

func isValid(num int) bool {
	if num >= 0 && num <= 10000 {
		return true
	} else {
		return false
	}
}
