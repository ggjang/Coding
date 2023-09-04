package main

import "fmt"

func main() {
	fmt.Println(solution(150000))
	fmt.Println(solution(580000))
}

func solution(price int) int {
	percent := 0.0
	if price >= 500000 {
		percent = 20
	} else if price >= 300000 {
		percent = 10
	} else if price >= 100000 {
		percent = 5
	}

	return salePrice(float64(price), percent)
}

func salePrice(price, percent float64) int {
	return int(price - (price * percent / 100))
}
