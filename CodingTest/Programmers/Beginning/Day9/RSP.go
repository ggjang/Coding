package main

func main() {
	println(solution("2"))
	println(solution("205"))
}

func solution(rsp string) string {
	winnig_map := map[byte]byte{
		'2': '0',
		'0': '5',
		'5': '2',
	}
	result := []byte{}
	for _, v := range rsp {
		result = append(result, winnig_map[byte(v)])
	}
	return string(result)
}
