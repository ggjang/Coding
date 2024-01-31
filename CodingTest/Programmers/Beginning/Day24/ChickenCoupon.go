package main

func main() {
	solution(100)
	solution(1081)
	solution(1999)
	solution(1000000)
}

func solution(chicken int) int {
	result := 0

	for chicken >= 10 {
		result += chicken / 10
		chicken = chicken/10 + chicken%10

		//fmt.Printf("result:%d, chicken:%d\n", result, chicken)
	}

	return result
}
