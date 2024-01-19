package main

import (
	"strings"
)

func main() {
	solution([]string{"aya", "yee", "u", "maa", "wyeoo"})
	//solution([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"})
}

func solution(babbling []string) int {
	result := len(babbling)
	babbling_token := []string{"aya", "ye", "woo", "ma"}

	for _, v := range babbling {
		temp := v
		for _, token := range babbling_token {
			temp = strings.Replace(temp, token, " ", -1)
		}

		for _, char := range temp {
			if char != ' ' {
				result--
				break
			}
		}
	}

	return result
}
