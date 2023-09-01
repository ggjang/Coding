package main

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := solution(array)
	println(result)

	array = []int{89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	result = solution(array)
	println(result)
}

func solution(numbers []int) float64 {
	sum := 0.0
	for _, v := range numbers {
		sum += float64(v)
	}
	return sum / float64(len(numbers))
}
