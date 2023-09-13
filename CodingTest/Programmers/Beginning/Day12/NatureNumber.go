package main

func main() {
	println(solution("aAb1B2cC34oOp"))
	println(solution("1a2b3c4d123"))
}

func solution(my_string string) int {
	result := 0
	for _, v := range my_string {
		if byte(v) >= '1' && byte(v) <= '9' {
			result += int(v - '0')
		}
	}
	return result
}
