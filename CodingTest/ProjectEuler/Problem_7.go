package main

// problem link : https://euler.synap.co.kr/problem=7

// 10001번째의 소수를 구하세요
func main() {
	count := 0
	for i := 2; true; i++ {
		yak := 0
		for ii := 1; ii <= i; ii++ {
			if i%ii == 0 {
				yak++
				if yak > 2 {
					break
				}
			}
		}

		if yak == 2 {
			count++
			if count == 10001 {
				println(i)
				break
			}
		}
	}
}
