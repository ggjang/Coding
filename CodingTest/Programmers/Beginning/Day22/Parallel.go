package main

func main() {

	//[[1, 4], [9, 2], [3, 8], [11, 6]]	1
	//[[3, 5], [4, 1], [2, 4], [5, 10]]	0

	println(solution([][]int{{1, 4}, {9, 2}, {3, 8}, {11, 6}}))

	println(solution([][]int{{3, 5}, {4, 1}, {2, 4}, {5, 10}}))
}

func solution(dots [][]int) int {
	// (1,2) / (3,4)
	// (1,3) / (2,4)
	// (1,4) / (2,3)

	if getInclination(dots[0], dots[1]) == getInclination(dots[2], dots[3]) {
		return 1
	}

	if getInclination(dots[0], dots[2]) == getInclination(dots[1], dots[3]) {
		return 1
	}

	if getInclination(dots[0], dots[3]) == getInclination(dots[1], dots[2]) {
		return 1
	}

	return 0
}

func getInclination(dot0, dot1 []int) float32 {
	return float32(dot0[1]-dot1[1]) / float32(dot0[0]-dot1[0])
}
