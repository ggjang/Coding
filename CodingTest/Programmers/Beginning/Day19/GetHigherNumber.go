package main

func main() {
	println(solution([]int{149, 180, 192, 170}, 167))
	println(solution([]int{180, 120, 140}, 190))
}

func solution(array []int, height int) int {
	result := 0
	for _, v := range array {
		if v > height {
			result++
		}
	}
	return result
}
