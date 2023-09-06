package main

import "fmt"

// problem link : https://euler.synap.co.kr/problem=6

// 1부터 100까지 자연수에 대해 "합의 제곱"과 "제곱의 합"의 차이는 얼마입니까?
func main() {
	// 합의 제곱
	sum := 0
	// 제곱의 합
	expone_sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		expone_sum += i * i
	}
	sum_expone := sum * sum
	fmt.Printf("합의 제곱:%d, 제곱의 합:%d, 두 값의 차: %d", sum_expone, expone_sum, sum_expone-expone_sum)
}
