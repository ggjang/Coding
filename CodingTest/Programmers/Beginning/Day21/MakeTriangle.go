package main

func main() {
	println(solution([]int{1, 2}))
	println(solution([]int{3, 6}))
	println(solution([]int{1, 1}))
	println(solution([]int{11, 7}))
	println(solution([]int{2, 2}))
}

// 가장 긴 변의 길이는 다른 두 변의 길이의 합보다 작아야 합니다

func solution(sides []int) int {
	result := 0
	if sides[0] > sides[1] {
		// sides[0]이 가장 긴 경우
		for i := sides[0] - sides[1] + 1; i <= sides[0]; i++ {
			result++
		}

		// 나오지 않은 다른 변이 가장 긴 경우
		for i := sides[0] + 1; i < sides[0]+sides[1]; i++ {
			result++
		}
	} else if sides[0] < sides[1] {
		// sides[1]이 가장 긴 경우
		for i := sides[1] - sides[0] + 1; i <= sides[1]; i++ {
			result++
		}

		// 나오지 않은 다른 변이 가장 긴 경우
		for i := sides[1] + 1; i < sides[0]+sides[1]; i++ {
			result++
		}
	} else {
		// 둘이 같고 두 변의 길이가 가장 긴 경우
		for i := 1; i < sides[0]; i++ {
			result++
		}

		// 나오지 않은 다른 변이 가장 긴 경우
		for i := sides[0]; i < sides[0]+sides[1]; i++ {
			result++
		}
	}

	return result
}
