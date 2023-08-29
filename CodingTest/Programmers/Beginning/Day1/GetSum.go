package main

func main() {
	result := solution(2, 3)
	println(result)

	result = solution(100, 2)
	println(result)
}

func solution(num1 int, num2 int) int {

	if isValid(num1) && isValid(num2) {
		return num1 + num2
	} else {
		return -1
	}
}

func isValid(num int) bool {
	if num >= -50000 && num <= 50000 {
		return true
	} else {
		return false
	}
}
