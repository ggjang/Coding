package main

import "fmt"

func main() {
	fmt.Println(solution(79))
	fmt.Println(solution(91))
	fmt.Println(solution(180))
}

func solution(angle int) int {
	if angle == 180 {
		return 4
	} else if angle < 180 && angle > 90 {
		return 3
	} else if angle == 90 {
		return 2
	} else {
		return 1
	}
}
