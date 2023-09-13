package main

import "fmt"

func main() {
	fmt.Println(solution("hi12392"))
	fmt.Println(solution("p2o4i8gj2"))
	fmt.Println(solution("abcde0"))
}

func solution(my_string string) []int {
	result := []int{}
	for _, v := range my_string {
		if byte(v) >= '0' && byte(v) <= '9' {
			result = append(result, int(v-'0'))
		}
	}

	for i := 0; i < len(result); i++ {
		for j := i; j < len(result); j++ {
			if result[i] > result[j] {
				temp := result[i]
				result[i] = result[j]
				result[j] = temp
			}
		}
	}

	return result
}
