package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := 0
	for i := 100; i < 1000; i++ {
		for ii := i + 1; ii < 1000; ii++ {
			num_to_str := strconv.Itoa(i * ii)
			temp := make([]byte, len(num_to_str))
			for k, _ := range num_to_str {
				temp[k] = num_to_str[len(num_to_str)-k-1]
			}
			if num_to_str == string(temp) {
				if i*ii > result {
					result = i * ii
					fmt.Println(result)
				}
			}
		}
	}
}
