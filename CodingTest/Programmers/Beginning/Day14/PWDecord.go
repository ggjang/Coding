package main

func main() {
	println(solution("dfjardstddetckdaccccdegk", 4))
	println(solution2("dfjardstddetckdaccccdegk", 4))
}

func solution(cipher string, code int) string {
	res := []rune{}
	for k, v := range cipher {
		if (k+1)%code == 0 {
			res = append(res, v)
		}
	}
	return string(res)
}

func solution2(cipher string, code int) string {
	res := []byte{}
	for i := code - 1; i < len(cipher); i += code {
		res = append(res, cipher[i])
	}
	return string(res)
}
