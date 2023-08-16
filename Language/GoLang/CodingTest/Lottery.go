package main

import "fmt"

func main() {
	lottos := []int{6, 14, 28, 5, 0, 0}
	win_nums := []int{6, 14, 28, 5, 34, 1}
	result := solution(lottos, win_nums)

	fmt.Println(result)
}

func solution(lottos []int, win_nums []int) [2]int {
	//Lotto rank
	// match 6 = 1
	// match 5 = 2
	// match 4 = 3
	// match 3 = 4
	// match 2 = 5
	// match 1 = 6
	const MAX_COUNT int = 6

	var rank [2]int

	var lottos_num int
	match_count := 0
	zero_count := 0

	for i := 0; i < MAX_COUNT; i++ {
		lottos_num = lottos[i]

		if lottos_num == 0 {
			zero_count++
		} else {
			for j := 0; j < MAX_COUNT; j++ {
				if lottos_num == win_nums[j] {
					fmt.Printf("Number is match : %d\n", lottos_num)
					match_count++
					break
				}
			}
		}
	}
	fmt.Printf("Zero : %d, Match : %d\n", zero_count, match_count)
	//Hightst rank
	switch match_count + zero_count {
	case 6:
		rank[0] = 1
	case 5:
		rank[0] = 2
	case 4:
		rank[0] = 3
	case 3:
		rank[0] = 4
	case 2:
		rank[0] = 5
	default:
		rank[0] = 6
	}

	//Lowest rank
	switch match_count {
	case 6:
		rank[1] = 1
	case 5:
		rank[1] = 2
	case 4:
		rank[1] = 3
	case 3:
		rank[1] = 4
	case 2:
		rank[1] = 5
	default:
		rank[1] = 6
	}

	return rank
}
