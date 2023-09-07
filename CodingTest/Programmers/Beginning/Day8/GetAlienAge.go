package main

func main() {
	result := solution(100)
	println(result)
}

func solution(age int) string {
	age_array := []int{}
	if age == 1000 {
		thousand := age / 1000
		hundrad := age/100 - (thousand * 10)
		ten := age/10 - (thousand * 100) - (hundrad * 10)
		one := age % 10
		age_array = []int{thousand, hundrad, ten, one}
	} else if age >= 100 && age < 1000 {
		hundrad := age / 100
		ten := age/10 - (hundrad * 10)
		one := age % 10
		age_array = []int{hundrad, ten, one}
	} else if age >= 10 && age < 100 {
		ten := age / 10
		one := age % 10
		age_array = []int{ten, one}
	} else {
		age_array = []int{age}
	}

	result := []byte{}
	for _, v := range age_array {
		result = append(result, byte(v+97))
	}

	return string(result)
}
