package main

import (
	"fmt"
	"time"
)

func main() {
	solution(10000000)
	//fmt.Println(result)

	goodsolution(10000000)
	//fmt.Println(result)
}

func solution(n int) []int {

	if isValid(n) {
		start := time.Now()
		result := []int{}
		for i := 1; i <= n; i++ {
			if i%2 != 0 {
				result = append(result, i)
			}
		}
		end := time.Since(start)
		fmt.Println("solution: ", end)
		return result
	} else {
		return nil
	}
}

func goodsolution(n int) []int {
	if isValid(n) {
		start := time.Now()
		res := []int{}
		for i := 1; i <= n; i += 2 {
			res = append(res, i)
		}
		end := time.Since(start)
		fmt.Println("goodsolution: ", end)
		return res
	} else {
		return nil
	}
}

func isValid(num int) bool {
	if num >= 1 && num <= 100 {
		return true
	}
	return true
}
