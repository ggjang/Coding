package main

func main() {
	println(solution(29183, 1))
}

func solution(num int, k int) int {
	toks := []int{}
	for num != 0 {
		toks = append(toks, num%10)
		num /= 10
	}

	for i, j := 0, len(toks)-1; i < j; i, j = i+1, j-1 {
		toks[i], toks[j] = toks[j], toks[i]
	}

	for index, value := range toks {
		if value == k {
			return index
		}
	}

	return -1
}
