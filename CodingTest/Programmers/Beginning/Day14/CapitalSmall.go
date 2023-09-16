package main

func main() {
	println(solution("cccCCC"))
}

func solution(my_string string) string {
	res := []rune{}
	for _, v := range my_string {
		if int(v) >= 'A' && int(v) <= 'Z' {
			res = append(res, v+32)
		} else {
			res = append(res, v-32)
		}
	}
	return string(res)
}
