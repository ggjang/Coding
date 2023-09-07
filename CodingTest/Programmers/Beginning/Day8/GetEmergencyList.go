package main

import "fmt"

func main() {
	//[3, 76, 24]	[3, 1, 2]
	fmt.Println(solution([]int{3, 76, 24}))
}

func solution(emergency []int) []int {

	result_map := make(map[int]int)
	for k, v := range emergency {
		result_map[v] = k
	}

	for i := 0; i < len(emergency)-1; i++ {
		for ii := i + 1; ii < len(emergency); ii++ {
			if emergency[i] < emergency[ii] {
				temp := emergency[i]
				emergency[i] = emergency[ii]
				emergency[ii] = temp
			}
		}
	}

	result := make([]int, len(emergency))
	for k, v := range emergency {
		//		result = append(result, result_map[v]+1)
		result[result_map[v]] = k + 1
	}

	return result
}
