package main

func main() {
	println(solution([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"}))
	println(solution([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"}))
	println(solution([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"}))
}

func solution(spell []string, dic []string) int {
	result := 2
	spell_map := make(map[string]bool)
	for _, v := range spell {
		spell_map[v] = false
	}

	for _, value := range dic {
		duplicateflag := false
		for _, v := range value {
			check, ok := spell_map[string(v)]
			if ok {
				if check {
					duplicateflag = true
					break
				} else {
					spell_map[string(v)] = true
				}
			}
		}

		if !duplicateflag {
			result = 1
			for _, v := range spell_map {
				if !v {
					result = 2
					break
				}
			}
			if result != 2 {
				return 1
			}
		}

		for k, _ := range spell_map {
			spell_map[k] = false
		}
	}
	return result
}
