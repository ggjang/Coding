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
