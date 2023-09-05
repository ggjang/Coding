package main

func main() {
	result := solution("jaron")
	println(result)
	result = solution("bread")
	println(result)
}

func solution(my_string string) string {
	result := []byte{}
	for i := len(my_string) - 1; i >= 0; i-- {
		result = append(result, my_string[i])
	}
	return string(result)
}
