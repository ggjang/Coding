package main

func main() {
	println(solution("abcabcadc"))
	println(solution("abdc"))
}

func solution(s string) string {
	toks := make(map[rune]int)
	for _, v := range s {
		value, exist := toks[v]
		if exist {
			toks[v] = value + 1
		} else {
			toks[v] = 1
		}
	}

	result := []rune{}
	for k, v := range toks {
		if v == 1 {
			result = append(result, k)
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

	return string(result)
}
