package main

import "fmt"

func main() {
	fmt.Println(solution(12))
	fmt.Println(solution(17))
	fmt.Println(solution(420))
}

func solution(n int) []int {
	temp := []int{}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			temp = append(temp, i)
		}
	}

	return makeSliceUnique(temp)
}

func makeSliceUnique(s []int) []int {
	keys := make(map[int]struct{})
	res := make([]int, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}
