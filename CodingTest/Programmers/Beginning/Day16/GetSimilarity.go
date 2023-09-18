package main

func main() {
	println(solution([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"}))
	println(solution([]string{"n", "omg"}, []string{"m", "dot"}))
}

func solution(s1 []string, s2 []string) int {
	toks := make(map[string]bool)
	for _, v := range s1 {
		_, ok := toks[v]
		if !ok {
			toks[v] = true
		}
	}

	result := 0

	for _, v := range s2 {
		_, ok := toks[v]
		if ok {
			result++
		}
	}
	return result
}
