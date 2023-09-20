package main

func main() {
	println(solution("Bcad"))
	println(solution("heLLo"))
	println(solution("Python"))
}

func solution(my_string string) string {
	result := []rune{}
	for _, v := range my_string {
		if v >= 'A' && v <= 'Z' {
			result = append(result, v+32)
		} else {
			result = append(result, v)
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
