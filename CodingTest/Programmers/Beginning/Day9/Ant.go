package main

func main() {
	println(solution(23))
	println(solution(24))
	println(solution(999))
}

func solution(hp int) int {
	result := 0
	result += hp / 5
	hp = hp % 5
	result += hp / 3
	hp = hp % 3
	result += hp

	return result
}
