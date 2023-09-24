package main

func main() {
	println(solution([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}))
}

func solution(dots [][]int) int {
	biggest_x := -256
	smallest_x := 256
	biggest_y := -256
	smallest_y := 256
	for _, v := range dots {
		if v[0] > biggest_x {
			biggest_x = v[0]
		}
		if v[1] > biggest_y {
			biggest_y = v[1]
		}

		if v[0] < smallest_x {
			smallest_x = v[0]
		}

		if v[1] < smallest_y {
			smallest_y = v[1]
		}

	}
	//println(biggest_x, smallest_x, biggest_x, smallest_y)

	return (biggest_x - smallest_x) * (biggest_y - smallest_y)
}
