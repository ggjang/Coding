package main

func main() {
	solution([]int{2, 4, 8})
}

func solution(common []int) int {
	// is it same multi?
	var multi int
	multi_bak := 0
	multi_flag := true
	for i := 1; i < len(common); i++ {
		if common[i-1] != 0 && common[i]%common[i-1] == 0 {
			multi = common[i] / common[i-1]
			if multi_bak == 0 {
				multi_bak = multi
			}
			if multi_bak != multi {
				multi_flag = false
				break
			}
		} else {
			multi_flag = false
			break
		}
	}

	if multi_flag {
		return common[len(common)-1] * multi
	}

	// is it same sub?
	var sub int
	sub_flag := true
	sub = common[1] - common[0]
	for i := 2; i < len(common); i++ {
		if common[i]-common[i-1] != sub {
			sub_flag = false
			break
		}
	}
	if sub_flag {
		return common[len(common)-1] + sub
	}

	return 0
}
