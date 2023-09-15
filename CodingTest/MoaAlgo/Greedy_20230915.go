package main

func main() {
	println(solution(4, 5, []int{1, 0, 3, 1, 2}, []int{0, 3, 0, 4, 0}))
	println(solution(2, 7, []int{1, 0, 2, 0, 1, 0, 2}, []int{0, 2, 0, 1, 0, 2, 0}))
	println(solution(6, 3, []int{50, 0, 31}, []int{10, 2, 45}))
}

// Algorithm
// 1. 배달 혹은 수거 물량이 있는 가장 멀리 있는 집 부터 배달 혹은 수거를 진행한다.
// 2. 최대 적재량 만큼 배달한다.
// 3. 최대 적재량 만큼 수거한다.
// 4. 배달 완료 후 수거를 진행한다.
// 5. 배달 혹은 수거 물량이 배달 혹은 수거 용량보다 많은 경우 남아있는 배달, 수거 용량 만큼만 처리하고 돌아온다.

func solution(cap int, n int, deliveries []int, pickups []int) int64 {
	// 트럭의 이동 거리
	res := int64(0)

	// 배달 용량
	del_cap := 0
	// 수거 용량
	pick_cap := 0

	for len(deliveries) != 0 || len(pickups) != 0 {
		del_top := len(deliveries) - 1
		pick_top := len(pickups) - 1

		var del int = 0
		if del_top >= 0 {
			del = deliveries[del_top]
		}

		var pickup int = 0
		if pick_top >= 0 {
			pickup = pickups[pick_top]
		}

		// 배달 용량이 남았거나 수거 용량이 남았음
		if del != 0 || pickup != 0 {
			if del_top > pick_top {
				// 배달 용량이 남은 집이 더 먼 경우
				res += int64(del_top+1) * 2
			} else if pick_top > del_top {
				// 수거 용량이 남은 집이 더 먼 경우
				res += int64(pick_top+1) * 2
			} else {
				// 배달, 수거 용량이 같은 경우
				res += int64(del_top+1) * 2
			}
		}

		// 먼 집 부터 배달 시작
		for true {
			top := len(deliveries) - 1
			if top < 0 {
				break
			}
			del := deliveries[top]
			// 남은 배달 용량보다 배달 해야할 양이 많은 경우
			// 배달 가능한 만큼 배달하고 다시 배달하러 온다.
			if del_cap+del > cap {
				deliveries[top] = del - (cap - del_cap)
				del_cap = 0
				break
			}
			del_cap += del
			deliveries = deliveries[:top]
		}

		// 먼 집 부터 수거 시작
		for true {
			top := len(pickups) - 1
			if top < 0 {
				break
			}
			pickup := pickups[top]
			// 남은 수거 용량보다 수거 해야할 양이 많은 경우
			// 수거 가능한 만큼 수거 하고 다시 수거하러 온다.
			if pick_cap+pickup > cap {
				pickups[top] = pickup - (cap - pick_cap)
				pick_cap = 0
				break
			}
			pick_cap += pickup
			pickups = pickups[:top]
		}
	}

	return res
}
