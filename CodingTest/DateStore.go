package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	term := []string{"A 6", "B 12", "C 3"}
	privacies := []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"}
	result := solution("2022.05.19", term, privacies)
	fmt.Println(result)
	term_2 := []string{"Z 3", "D 5"}
	privacies_2 := []string{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"}
	result = solution("2020.01.01", term_2, privacies_2)
	fmt.Println(result)
}

func solution(today string, terms []string, privacies []string) []int {
	// 오늘 날짜 일수
	today_atday := 0

	// 약관 기간 map 일수로 변환
	terms_map := make(map[string]int)

	// 결과
	var result []int

	// today 형태
	// today를 모두 일수로 변환
	// 1개월은 28일
	// 1년은
	toks := strings.Split(today, ".")
	if len(toks) != 3 {
		return nil
	} else {
		y_date, _ := strconv.Atoi(toks[0])
		m_date, _ := strconv.Atoi(toks[1])
		d_date, _ := strconv.Atoi(toks[2])
		today_atday += y_date * 12 * 28
		today_atday += m_date * 28
		today_atday += d_date
	}

	// 약관 A, B, C ...
	// 약관 제한사항 확인 terms는 최대 20 1 <= terms <= 20
	terms_len := len(terms)
	if terms_len >= 1 && terms_len <= 20 {
		for _, v := range terms {
			toks := strings.Split(v, " ")
			term, _ := strconv.Atoi(toks[1])
			terms_map[toks[0]] = term * 28
		}
	} else {
		return nil
	}

	// privacies limit check
	privacies_len := len(privacies)
	if privacies_len >= 1 && privacies_len <= 100 {

	} else {
		return nil
	}
	for k, v := range privacies {
		toks_privacies := strings.Split(v, " ")
		toks_date := strings.Split(toks_privacies[0], ".")
		toks_term := toks_privacies[1]
		y_date, _ := strconv.Atoi(toks_date[0])
		m_date, _ := strconv.Atoi(toks_date[1])
		d_date, _ := strconv.Atoi(toks_date[2])
		date := (y_date * 12 * 28) + (m_date * 28) + d_date

		term := terms_map[toks_term]
		date_diff := today_atday - date
		if date_diff >= term {
			result = append(result, k+1)
		}
	}

	return result
}
