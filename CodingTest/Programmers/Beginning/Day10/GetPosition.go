package main

func main() {
	println(solution([]int{2, 4}))
	println(solution([]int{-7, 9}))
}

func solution(dot []int) int {
	// this case will be 1 or 3
	if dot[0]*dot[1] > 0 {
		if dot[0] > 0 {
			return 1
		} else {
			return 3
		}
	} else { // this case whill be 2 or 4
		if dot[0] > 0 {
			return 4
		} else {
			return 2
		}
	}
}
