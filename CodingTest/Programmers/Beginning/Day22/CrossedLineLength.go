package main

func main() {
	println(solution([][]int{{0, 1}, {2, 5}, {3, 9}}))
}

func solution(lines [][]int) int {
	// Use imos algorithm
	all_line := make([]int, 201)

	for _, v := range lines {
		all_line[v[0]+100]++
		all_line[v[1]+100]--
	}

	var ans int
	if all_line[0] > 1 {
		ans = 1
	} else {
		ans = 0
	}
	for i := 1; i < 201; i++ {
		all_line[i] += all_line[i-1]
		if all_line[i] > 1 {
			ans++
		}
	}

	return ans
}
