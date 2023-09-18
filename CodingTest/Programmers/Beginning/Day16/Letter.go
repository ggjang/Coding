package main

func main() {
	println(solution("happy birthday!"))
	println(solution("I love you~"))
}

func solution(message string) int {
	return len(message) * 2
}
