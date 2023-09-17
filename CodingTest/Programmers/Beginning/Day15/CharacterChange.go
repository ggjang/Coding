package main

func main() {
	println(solution("hello", 1, 3))
}

func solution(my_string string, num1 int, num2 int) string {
	result := []rune{}
	for _, v := range my_string {
		result = append(result, v)
	}
	temp := result[num1]
	result[num1] = result[num2]
	result[num2] = temp
	return string(result)
}
