package main

func main() {
	result := solution(7)
	println(result)

	result = solution(10)
	println(result)

	result = solution(1)
	println(result)
}

func solution(n int) int {
	if isValid(n) {
		gcd := -1
		for i := 1; i <= n && i <= 6; i++ {
			if n%i == 0 && 6%i == 0 {
				gcd = i
			}
		}
		lcm := (n * 6) / gcd
		return lcm / 6
	}
	return 0
}

func isValid(n int) bool {
	if n >= 1 && n <= 100 {
		return true
	}
	return false
}
