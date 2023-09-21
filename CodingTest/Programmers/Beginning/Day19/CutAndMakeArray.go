package main

import "fmt"

func main() {
	fmt.Println(solution("abc1Addfggg4556b", 6))
	fmt.Println(solution("abcdef123", 3))
}

func solution(my_str string, n int) []string {
	result := []string{}
	result_str := ""
	for i := 0; i < len(my_str); i++ {
		if i%n == 0 && i != 0 {
			result = append(result, result_str)
			result_str = ""
		}
		result_str += string(my_str[i])
	}
	result = append(result, result_str)

	return result
}
