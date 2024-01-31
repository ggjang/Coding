package main

func main() {

}

func solution(num int, total int) []int {
	var result []int
	start_number := 0
	for true {
		temp_start_number := start_number
		temp_sum := 0
		for count := 0; count < num; count++ {
			temp_sum += temp_start_number
			temp_start_number++
		}

		if temp_sum == total {
			break
		} else if temp_sum < total {
			start_number++
		} else {
			start_number--
		}
	}
	for i := 0; i < num; i++ {
		result = append(result, start_number)
		start_number++
	}

	return result
}
