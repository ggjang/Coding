package main

import "fmt"

func main() {
	println(solution([][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}}))
}

func solution(board [][]int) int {
	result := len(board) * len(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				fmt.Printf("is zero case %d,%d\n", i, j)
				// check
				// (i-1,j+1) (i, j+1) (i+1,j+1)
				// (i-1,j)   (i, j)   (i+1,j)
				// (i-1,j-1) (i, j-1) (i+1,j-1)
				safezone_flag := true
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						temp_x := x
						temp_y := y
						if x < 0 {
							temp_x = 0
						}
						if y < 0 {
							temp_y = 0
						}
						if x > len(board)-1 {
							temp_x = len(board) - 1
						}
						if y > len(board)-1 {
							temp_y = len(board) - 1
						}

						if board[temp_x][temp_y] == 1 {
							fmt.Printf("%d,%d,%d,%d\n", i, j, temp_x, temp_y)
							result--
							safezone_flag = false
							break
						}
					}
					if safezone_flag == false {
						break
					}
				}

			} else {
				fmt.Printf("is not zero case %d,%d\n", i, j)
				result--
			}
		}
	}
	return result
}
