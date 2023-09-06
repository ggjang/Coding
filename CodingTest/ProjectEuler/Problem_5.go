package main

import "fmt"

// problem link https://euler.synap.co.kr/problem=5

// 1 ~ 20 사이의 어떤 수로도 나누어 떨어지는 가장 작은수는 얼마입니까?
func main() {
	escape := true
	for i := 20; escape; i++ {
		var div int
		for div = 1; div <= 20; div++ {
			if i%div != 0 {
				break
			}
		}
		if div > 20 {
			escape = false
			fmt.Println(i)
		}
	}
}
