package main

import (
	"fmt"
	"strconv"
)

func main() {
	solution("10", "11")
	solution("101", "111")
}

func solution(bin1 string, bin2 string) string {
	dec1 := makeDec(bin1)
	dec2 := makeDec(bin2)

	fmt.Printf("bin1(%s) -> dec1(%d)\n", bin1, dec1)
	fmt.Printf("bin2(%s) -> dec2(%d)\n", bin2, dec2)

	dec_sum := dec1 + dec2

	fmt.Printf("dec_sum(%d) -> bin_sum(%b)\n", dec_sum, dec_sum)
	result := strconv.FormatInt(int64(dec_sum), 2)

	return result
}

func makeDec(bin string) int {
	result := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			result += square(2, len(bin)-i)
		}
	}
	return result
}

func square(num, sqr int) int {
	result := 1
	for i := 1; i < sqr; i++ {
		result *= 2
	}
	return result
}
