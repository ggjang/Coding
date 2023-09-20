package main

func main() {
	println(solution(144))
	println(solution(976))
}

func solution(n int) int {
	measure := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			measure = append(measure, i)
		}
	}
	for _, v := range measure {
		if v*v == n {
			return 1
		}
	}

	return 2
}
