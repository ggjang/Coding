package main

import "strings"

func main() {
	result := usePkgSolution("abcdef", "f")
	println(result)
	result = solution("abcdef", "f")
	println(result)
}

func usePkgSolution(my_string string, letter string) string {
	result := strings.Replace(my_string, letter, "", -1)
	return result
}

func solution(my_string string, letter string) string {
	var result []byte
	for _, v := range my_string {
		if byte(v) != letter[0] {
			result = append(result, byte(v))
		}
	}

	return string(result)
}
