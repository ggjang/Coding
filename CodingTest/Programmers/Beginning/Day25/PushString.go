package main

import "fmt"

func main() {
	fmt.Println(solution("hello", "ohell"))
}

func solution(A string, B string) int {
	if A == B {
		return 0
	} else {
		for i := 0; i < len(A); i++ {
			pushed := A[len(A)-i:]
			pushed += A[:len(A)-i]
			if pushed == B {
				return i
			}
			//fmt.Println(pushed)
		}
	}

	return -1
}
