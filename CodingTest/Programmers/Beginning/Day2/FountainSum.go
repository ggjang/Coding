package main

import "fmt"

func main() {
	result := solution(1, 2, 3, 4)
	fmt.Println(result)
	result = solution(9, 2, 1, 3)
	fmt.Println(result)
}

func solution(numer1 int, denom1 int, numer2 int, denom2 int) []int {
	if isValid(numer1) && isValid(numer2) && isValid(denom1) && isValid(denom2) {
		var result []int

		numer := (numer1 * denom2) + (numer2 * denom1)
		denom := denom1 * denom2

		var gcd int
		for i := 1; i <= numer && i <= denom; i++ {
			if numer%i == 0 && denom%i == 0 {
				gcd = i
			}
		}

		result = append(result, numer/gcd)
		result = append(result, denom/gcd)
		return result
	} else {
		return nil
	}
}

func isValid(num int) bool {
	if num > 0 && num < 1000 {
		return true
	} else {
		return false
	}
}
