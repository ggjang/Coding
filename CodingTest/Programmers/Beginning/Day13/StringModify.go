package main

import "fmt"

func main() {
	fmt.Println(solution("people"))
}

func solution(my_string string) string {
	toks_map := make(map[rune]bool)
	result := []rune{}

	for _, v := range my_string {
		if _, exist := toks_map[v]; !exist {
			result = append(result, v)
			toks_map[v] = true
		}
	}

	return string(result)
}
