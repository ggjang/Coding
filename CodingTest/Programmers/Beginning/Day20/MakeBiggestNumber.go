package main

func main() {
	println(solution([]int{-2, 2}))
}

func solution(numbers []int) int {
	result := -100000000
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if result < numbers[i]*numbers[j] {
				result = numbers[i] * numbers[j]
			}
		}
	}
	return result
}
