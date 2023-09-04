package main

import (
	"fmt"
	"time"
)

// problem link https://euler.synap.co.kr/problem=3

// 가장 큰 소인수 구하기
func main() {
	input := 600851475143
	c_org := make(chan time.Duration, 1)
	c_half := make(chan time.Duration, 1)
	c_root := make(chan time.Duration, 1)

	go solution_org(input, c_org)
	go solution_half(input, c_half)
	go solution_root(input, c_root)

	for true {
		select {
		case res := <-c_org:
			fmt.Println("solution_org Time:", res)
		case res := <-c_half:
			fmt.Println("solution_half Time:", res)
		case res := <-c_root:
			fmt.Println("solution_root Time:", res)
		case <-time.After(time.Second * 20):
			fmt.Println("It's Time Over (60sec)")
			return
		}
	}
}

// 소수 Study
// 2 ~ 판별하는 수 전까지 나눠보고 나머지가 0이 안나온다면 소수로 정의
func solution_org(input int, c chan time.Duration) {
	start := time.Now()
	for i := 2; i <= input; i++ {
		if input%i == 0 {
			ok := true
			for ii := 2; ii < i; ii++ {
				if i%ii == 0 {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("%d is ok\n", i)
			}
		}
	}
	end := time.Since(start)
	c <- end
}

// 위 방법과 동일하나 절반만 확인한다.
func solution_half(input int, c chan time.Duration) {
	start := time.Now()
	for i := 2; i <= input/2; i++ {
		if input%i == 0 {
			ok := true
			for ii := 2; ii < i/2; ii++ {
				if i%ii == 0 {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("%d is ok\n", i)
			}
		}
	}
	end := time.Since(start)
	c <- end
}

// 해당 숫자의 루트 N 까지 확인하는 방법

func solution_root(input int, c chan time.Duration) {
	start := time.Now()
	for i := 2; i*i <= input; i++ {
		if input%i == 0 {
			ok := true
			for ii := 2; ii*ii < i; ii++ {
				if i%ii == 0 {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("%d is ok\n", i)
			}
		}
	}
	c <- time.Since(start)
}
