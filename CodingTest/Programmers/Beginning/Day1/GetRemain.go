package main

func main() {
	result := solution(3, 2)
	println(result)
	result = solution(10, 5)
	println(result)
}

func solution(num1 int, num2 int) int {
	if isValid(num1) && isValid(num2) {
		return num1 % num2
	} else {
		return -1
	}
}

func isValid(num int) bool {
	if num > 0 && num <= 100 {
		return true
	} else {
		return false
	}
}
