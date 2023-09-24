package main

func main() {

}

func solution(keyinput []string, board []int) []int {

	x_max_pos := board[0] / 2
	x_min_pos := -board[0] / 2
	y_max_pos := board[1] / 2
	y_min_pos := -board[1] / 2

	default_pos := []int{0, 0}
	for _, v := range keyinput {

		if v == "right" {
			default_pos[0]++
			if default_pos[0] > x_max_pos {
				default_pos[0] = x_max_pos
			}
		} else if v == "left" {
			default_pos[0]--
			if default_pos[0] < x_min_pos {
				default_pos[0] = x_min_pos
			}
		} else if v == "up" {
			default_pos[1]++
			if default_pos[1] < y_max_pos {
				default_pos[1] = y_max_pos
			}
		} else if v == "down" {
			default_pos[1]--
			if default_pos[1] < y_min_pos {
				default_pos[1] = y_min_pos
			}
		}

	}

	return default_pos
}
