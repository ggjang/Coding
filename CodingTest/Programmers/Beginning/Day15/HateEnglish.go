package main

import "strconv"

func main() {
	println(solution("onetwothreefourfivesixseveneightnine"))
}

func solution(numbers string) int64 {
	// one two three foure five six seven eigth nine zero
	// o : one (unique)
	// t : two, three (need check)
	// f : four, five (need check)
	// s : six, seven (need check)
	// e : eigth (unique)
	// n : nine (unique)
	// z : zero (unique)

	res := []byte{}
	for i := 0; i < len(numbers); i++ {
		switch numbers[i] {
		// first check unique case
		case 'o':
			res = append(res, '1')
			i += 2
		case 'e':
			res = append(res, '8')
			i += 4
		case 'n':
			res = append(res, '9')
			i += 3
		case 'z':
			res = append(res, '0')
			i += 3
		// need check
		case 't':
			if numbers[i+1] == 'w' {
				res = append(res, '2')
				i += 2
			} else {
				res = append(res, '3')
				i += 4
			}
		case 'f':
			if numbers[i+1] == 'o' {
				res = append(res, '4')
				i += 3
			} else {
				res = append(res, '5')
				i += 3
			}
		case 's':
			if numbers[i+1] == 'i' {
				res = append(res, '6')
				i += 2
			} else {
				res = append(res, '7')
				i += 4
			}
		}
	}

	result, _ := strconv.Atoi(string(res))

	return int64(result)
}
