package main

import "fmt"

func main() {
	input := 600851475143
	for i := 2; i <= input; i++ {
		if input%i == 0 {
			ok := true
			for ii := 2; ii < i; ii++ {
				if i%ii == 0 {
					ok = false
					break
				}
			}
			if ok {
				fmt.Print(("Yes\n"))
			} else {
				fmt.Print(("No\n"))
			}
		}
	}
}
