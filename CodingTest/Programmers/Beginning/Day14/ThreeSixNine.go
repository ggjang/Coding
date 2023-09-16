package main

func main() {
	println(solution(3))
	println(solution2(3))
	println(solution(29423))
	println(solution2(29423))
}

// my solution
func solution(order int) int {
	res_list := []int{}
	for i := 10; order != 0; i *= 10 {
		value := order % i
		res_list = append(res_list, value)
		if len(res_list) == 1 {
			order -= res_list[len(res_list)-1]
			continue
		}
		res_list[len(res_list)-1] = res_list[len(res_list)-1] / (i / 10)
		order -= res_list[len(res_list)-1] * (i / 10)
	}

	res := 0
	for _, v := range res_list {
		if v == 3 || v == 6 || v == 9 {
			res++
		}
	}
	return res
}

// good solution
func solution2(order int) int {
	res := 0
	for order > 0 {
		if digit := order % 10; digit%3 == 0 && digit > 0 {
			res++
		}
		order /= 10
	}
	return res
}
