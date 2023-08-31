package main

func main() {
	array := []int{1, 2, 3, 3, 3, 4}
	result := solution(array)
	println(result)

	array = []int{1, 1, 2, 2}
	result = solution(array)
	println(result)

	array = []int{1}
	result = solution(array)
	println(result)
}

func solution(array []int) int {

	if isValid(array) {
		number_map := make(map[int]int)
		for _, v := range array {
			val := number_map[v]
			number_map[v] = val + 1
		}
		var result int

		for _, v1 := range number_map {
			for k2, v2 := range number_map {
				if v1 > v2 {
					delete(number_map, k2)
				}
			}
		}

		if len(number_map) > 1 {
			return -1
		} else {
			for k := range number_map {
				return k
			}
		}

		return result
	} else {
		return 0
	}
}

func isValid(array []int) bool {
	if len(array) <= 0 || len(array) >= 100 {
		return false
	}

	for _, v := range array {
		if v < 0 || v >= 1000 {
			return false
		}
	}

	return true
}
