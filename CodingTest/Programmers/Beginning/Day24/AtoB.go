package main

func main() {
	println(solution("olleh", "hello"))
	println(solution("allpe", "apple"))
}

func solution(before string, after string) int {
	var before_toks []rune = []rune(before)
	flag := true

	for _, after_token := range after {
		flag = true
		for i, before_token := range before_toks {
			if after_token == before_token {
				before_toks[i] = ' '
				flag = false
				break
			}
		}
		if flag {
			return 0
		}
	}

	return 1
}
