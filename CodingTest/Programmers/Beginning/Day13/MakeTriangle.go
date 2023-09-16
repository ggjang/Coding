package main

func main() {
	println(solution([]int{1, 2, 3}))
}

func solution(sides []int) int {
	biggest_val := sides[0]
	biggest_key := 0
	for i := 1; i < len(sides); i++ {
		if biggest_val < sides[i] {
			biggest_val = sides[i]
			biggest_key = i
		}
	}

	sum := 0
	for k, v := range sides {
		if biggest_key != k {
			sum += v
		}
	}

	if sum > biggest_val {
		return 1
	} else {
		return 2
	}
}
