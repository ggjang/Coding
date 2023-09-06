package main

import (
	"fmt"
	"strconv"
)

func main() {

	test_count := 0
	_, err := fmt.Scanf("%d\n", &test_count)
	if err != nil {
		panic(err)
	}
	input := make([]uint32, test_count)
	for i := 0; i < test_count; i++ {
		_, err := fmt.Scanf("%d\n", &input[i])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("My Solution")
	mySolution(input)

	fmt.Printf("\nUse Shift\n")
	useShift(input)

}

func mySolution(input []uint32) {
	for _, v := range input {
		// dec -> hex
		// hex -> byte
		input_byte := []byte(fmt.Sprintf("%08x", v))

		// place change
		temp_byte := make([]byte, len(input_byte))
		index := 0
		for i := len(input_byte) - 1; i >= 0; i -= 2 {
			temp_byte[index] = input_byte[i-1]
			temp_byte[index+1] = input_byte[i]
			index += 2
		}

		// byte -> string(hex)
		// string(hex) -> dec
		number, _ := strconv.ParseInt(fmt.Sprintf("%s", temp_byte), 16, 64)
		fmt.Println(number)
	}
}

func useShift(input []uint32) {
	for _, v := range input {
		var result uint32
		result += (v >> 24) & 0xFF << 0
		result += (v >> 16) & 0xFF << 8
		result += (v >> 8) & 0xFF << 16
		result += (v >> 0) & 0xFF << 24
		fmt.Println(result)
	}
}
