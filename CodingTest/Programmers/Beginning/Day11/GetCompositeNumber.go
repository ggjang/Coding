package main

func main() {
	println(solution(10))
	println(solution(15))

}

func solution(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				temp++
			}
		}
		if temp >= 3 {
			res++
		}
	}

	return res
}
