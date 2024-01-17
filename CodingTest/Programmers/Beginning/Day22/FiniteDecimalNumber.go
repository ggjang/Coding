package main

import (
	"fmt"
)

func main() {
	println(solution(7, 20))
	println(solution(11, 22))
	println(solution(12, 21))
}

func solution(a, b int) int {
	if a > b {
		for i := 1; i <= b; i++ {
			if a%i == 0 && b%i == 0 {
				a /= i
				b /= i
			}
		}
		fmt.Printf("a:%d, b:%d\n", a, b)
	} else if a < b {
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				a /= i
				b /= i
			}
		}
		fmt.Printf("a:%d, b:%d\n", a, b)
	} else {
		return 1
	}

	for true {
		if b%2 == 0 {
			b /= 2
		} else if b%5 == 0 {
			b /= 5
		} else {
			break
		}
	}

	fmt.Printf("a:%d, b:%d\n", a, b)

	if b%2 == 0 || b%5 == 0 || b == 1 {
		return 1
	} else {
		return 2
	}
}
