package main

func main() {
	result := solution("hello", 3)
	println(result)
}

func solution(my_string string, n int) string {
	temp := []byte{}
	for k, _ := range my_string {
		for i := 0; i < n; i++ {
			temp = append(temp, my_string[k])
		}
	}
	return string(temp)
}
