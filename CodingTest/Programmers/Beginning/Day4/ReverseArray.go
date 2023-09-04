package main

func main() {

}

func solution(num_list []int) []int {
	length := len(num_list)
	result := make([]int, length)
	for k, v := range num_list {
		result[length-1-k] = v
	}
	return result
}
