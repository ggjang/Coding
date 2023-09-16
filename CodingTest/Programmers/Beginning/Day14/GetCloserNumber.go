package main

func main() {
	println(solution([]int{3, 10, 28}, 20))
	println(solution([]int{10, 11, 12}, 13))
	println(solution([]int{1, 2, 4, 1}, 3))
}

func solution(array []int, n int) int {
	// get diff at abs
	diff_list := []int{}
	for _, v := range array {
		diff := n - v
		if diff >= 0 {
			diff_list = append(diff_list, diff)
		} else if diff < 0 {
			diff_list = append(diff_list, -diff)
		}
	}

	// get smallest differnce from list
	smallest_diff := 100
	for _, v := range diff_list {
		if v < smallest_diff {
			smallest_diff = v
		}
	}

	// get number from array as list
	res_list := []int{}
	for k, v := range diff_list {
		if v == smallest_diff {
			res_list = append(res_list, array[k])
		}
	}

	// get result
	for i := 0; i < len(res_list); i++ {
		for j := i; j < len(res_list); j++ {
			if res_list[i] > res_list[j] {
				temp := res_list[i]
				res_list[i] = res_list[j]
				res_list[j] = temp
			}
		}
	}

	return res_list[0]
}
