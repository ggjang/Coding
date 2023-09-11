package main

import "strings"

func main() {
	println(solution(".... . .-.. .-.. ---"))
	println(solution(".--. -.-- - .... --- -."))
}

func solution(letter string) string {
	morse_raw := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morse := make(map[string]byte)
	for k, v := range morse_raw {
		morse[v] = byte('a' + k)
	}

	letter_toks := strings.Split(letter, " ")

	result := []byte{}
	for _, v := range letter_toks {
		result = append(result, morse[v])
	}

	return string(result)
}
