package main

func main() {
	solution([][]int{{80, 70}, {90, 50}, {40, 70}, {50, 80}})
	solution([][]int{{80, 70}, {70, 80}, {30, 50}, {90, 100}, {100, 90}, {100, 100}, {10, 30}})
}

func solution(score [][]int) []int {
	avg := make([]float64, len(score))
	sort_avg := make([]float64, len(score))
	result := make([]int, len(score))
	score_table := make(map[float64]int)

	//result := make([]int, len(score))

	for k, v := range score {
		avg[k] = (float64(v[0]) + float64(v[1])) / 2.0
	}

	copy(sort_avg, avg)
	// sort
	for i := 0; i < len(sort_avg); i++ {
		for j := i + 1; j < len(sort_avg); j++ {
			if sort_avg[i] < sort_avg[j] {
				temp := sort_avg[i]
				sort_avg[i] = sort_avg[j]
				sort_avg[j] = temp
			}
		}
	}

	for i := 0; i < len(sort_avg); i++ {
		if _, ok := score_table[sort_avg[i]]; !ok {
			score_table[sort_avg[i]] = i + 1
		}
	}

	for i := 0; i < len(avg); i++ {
		result[i] = score_table[avg[i]]
	}

	return result
}
