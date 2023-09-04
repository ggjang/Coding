package main

func main() {
	result := solution(7, 10)
	println(result)

	result = solution(4, 12)
	println(result)
}

func solution(slice int, n int) int {
	if isValid(slice, n) {

		result := 0
		for i := 1; true; i++ {
			if (i*slice)-n >= 0 {
				result = i
				break
			}
		}
		return result
	}
	return 0
}

func isValid(slice, n int) bool {
	if slice < 2 || slice > 10 {
		return false
	}
	if n < 1 || n > 100 {
		return false
	}
	return true
}
