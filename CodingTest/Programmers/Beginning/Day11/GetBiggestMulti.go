package main

func main() {
	println(solution([]int{1, 2, 3, 4, 5}))
	println(solution([]int{0, 31, 24, 10, 1, 9}))
}

func solution(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	return numbers[len(numbers)-1] * numbers[len(numbers)-2]
}
