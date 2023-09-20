package main

func main() {
	println(solution("abcd", "cd"))
}
func solution(str1 string, str2 string) int {

	for i := 0; i < len(str1)-len(str2)+1; i++ {
		if str1[i:i+len(str2)] == str2 {
			return 1
		}
	}
	return 2
}
