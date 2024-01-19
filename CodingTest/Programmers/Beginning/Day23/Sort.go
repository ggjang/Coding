package main

import (
	"fmt"
	"math"
)

func main() {
	//slolution([]int{1, 2, 3, 4, 5, 6}, 4)
	slolution([]int{10000, 20, 36, 47, 40, 6, 10, 7000}, 30)
}

func slolution(numlist []int, n int) []int {
	for i := 0; i < len(numlist); i++ {
		for j := i + 1; j < len(numlist); j++ {
			if numlist[i] > numlist[j] {
				temp := numlist[j]
				numlist[j] = numlist[i]
				numlist[i] = temp
			}
		}
	}
	fmt.Println(numlist)

	for i := 0; i < len(numlist); i++ {
		for j := i + 1; j < len(numlist); j++ {
			if math.Abs(float64(numlist[i]-n)) >= math.Abs(float64(numlist[j]-n)) {
				fmt.Printf("i:%d, j:%d\n", numlist[i], numlist[j])
				temp := numlist[j]
				numlist[j] = numlist[i]
				numlist[i] = temp
				if numlist[i] < numlist[j] {
					temp := numlist[j]
					numlist[j] = numlist[i]
					numlist[i] = temp
				}
				fmt.Printf("%v\n", numlist)
			}
		}
	}

	return numlist
}
