package main

func main() {
	println(solution([]string{"meosseugi", "1234"}, [][]string{{"rardss", "123"}, {"yyoom", "1234"}, {"meosseugi", "1234"}}))
	println(solution([]string{"programmer01", "15789"}, [][]string{{"programmer02", "111111"}, {"programmer00", "134"}, {"programmer01", "1145"}}))
	println(solution([]string{"rabbit04", "98761"}, [][]string{{"jaja11", "98761"}, {"krong0313", "29440"}, {"rabbit00", "111333"}}))
}

func solution(id_pw []string, db [][]string) string {
	for _, db_token := range db {
		if db_token[0] == id_pw[0] {
			if db_token[1] == id_pw[1] {
				return "login"
			} else {
				return "wrong pw"
			}
		}
	}

	return "fail"
}
