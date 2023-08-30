package main

// problem link https://euler.synap.co.kr/problem=1

// 1000보다 작은 자연수 중에서 3 또는 5의 배수를 모두 더하면
func main() {
	result := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
			println(result, " ", i)
		}
	}
}
