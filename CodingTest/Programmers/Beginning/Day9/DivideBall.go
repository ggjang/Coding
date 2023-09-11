package main

func main() {
	println(solution(30, 10))
	println(good_solution(30, 10))
}

func good_solution(balls int, share int) int {
	answer := 1
	for i := 0; i < share; i++ {
		answer *= balls - i
		answer /= i + 1
	}

	return answer
}

func solution(balls int, share int) int {
	//result := int(new_factorial(balls, share) / factorial(balls-share))
	return new_factorial(balls, share)
}

func new_factorial(num1 int, num2 int) int {
	numerator_toks := []int{}
	denominator_toks := []int{}

	for i := num2 + 1; i <= num1; i++ {
		numerator_toks = append(numerator_toks, i)
	}

	for i := 1; i <= num1-num2; i++ {
		denominator_toks = append(denominator_toks, i)
	}

	for numer_k, numer := range numerator_toks {
		for denom_k, denom := range denominator_toks {
			if numer%denom == 0 && denom != 1 {
				numerator_toks[numer_k] = numer / denom
				denominator_toks = append(denominator_toks[:denom_k], denominator_toks[denom_k+1:]...)

				break
			}
		}
	}

	numerator := 1
	for _, v := range numerator_toks {
		numerator *= v
	}

	denominator := 1
	for _, v := range denominator_toks {
		denominator *= v
	}
	if denominator == 0 {
		return numerator
	}
	return numerator / denominator
}
