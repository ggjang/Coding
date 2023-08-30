package main

// problem link https://euler.synap.co.kr/problem=2

// 피보나치 수열에서 4백만 이하이면서 짝수인 항의 합
func main() {
	result := 0
	n2 := 0
	n1 := 1
	for i := 2; true; i++ {
		n2, n1 = n1, n1+n2
		if n1 < 4000000 {
			if n1%2 == 0 {
				result += n1
			}
		} else {
			break
		}
	}

	println(result)
}
