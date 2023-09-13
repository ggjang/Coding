package main

func main() {
	println(solution("bus"))
	println(solution("nice to meet you"))
}

func solution(my_string string) string {
	str_byte := []byte{}
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range my_string {
		isvowel := false
		for _, vowel := range vowels {
			if byte(v) == vowel {
				isvowel = true
				break
			}
		}
		if !isvowel {
			str_byte = append(str_byte, byte(v))
		}
	}
	return string(str_byte)
}
