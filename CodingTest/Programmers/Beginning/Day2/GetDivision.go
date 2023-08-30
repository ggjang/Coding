package main

func main() {
	result := solution(3, 2)
	println(result)

	result = solution(7, 3)
	println(result)

	result = solution(1, 16)
	println(result)
}

func solution(num1 int, num2 int) int {

	if isValid(num1) && isValid(num2) {
		var result float64
		result = float64(num1) / float64(num2) * 1000.0
		return int(result)
	} else {
		return 0
	}
}

func isValid(num int) bool {
	if num > 0 && num <= 100 {
		return true
	} else {
		return false
	}
}
