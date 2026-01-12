package ds

import (
	"fmt"
)

func RollingMinMax(nums []int, k int) (mins []int, maxs []int, err error) {

	fmt.Println("-------------------------------------------------------------------------------------")

	maxQ := NewDeque(k) // Monotonic ordering (indexes)
	minQ := NewDeque(k)

	maxs = make([]int, 0, k)
	mins = make([]int, 0, k)

	k = min(len(nums), k)

	purge := func(Q *Deque, curr_idx int) {
		for !Q.Empty() {
			best_idx, _ := Q.PeekFront()
			if best_idx > curr_idx-k {
				break
			}
			front_idx, _ := Q.PopFront()
			fmt.Println("Purging idx:", front_idx, "value:", nums[front_idx])
		}
	}

	monotonic_push := func(Q *Deque, new_idx int, pop_cond func(old int, new int) bool) {
		new_value := nums[new_idx]

		for !Q.Empty() {
			back_idx, _ := Q.PeekBack()
			old_value := nums[back_idx]
			if pop_cond(old_value, new_value) {
				break
			}
			Q.PopBack()
			fmt.Println("PopBack: ", nums[back_idx])
		}

		Q.PushBack(new_idx)
		fmt.Println("Pushback:", new_value)
	}

	for curr_idx, num := range nums {
		fmt.Println("Processing idx:", curr_idx, "=>", num)

		// Drop any expired elements
		fmt.Println("Purging maxQ")
		purge(maxQ, curr_idx)

		fmt.Println("Purging minQ")
		purge(minQ, curr_idx)

		// Ensure monotonicity

		fmt.Println("Pushing to maxQ")
		monotonic_push(maxQ, curr_idx, func(old_value, new_value int) bool { return old_value > new_value })

		fmt.Println("Pushing to minQ")
		monotonic_push(minQ, curr_idx, func(old_value, new_value int) bool { return old_value < new_value })

		if curr_idx >= k-1 {
			max_idx, _ := maxQ.PeekFront()
			max_value := nums[max_idx]

			min_idx, _ := minQ.PeekFront()
			min_value := nums[min_idx]

			maxs = append(maxs, max_value)
			mins = append(mins, min_value)

			fmt.Println("Rolling max:", max_value, "min:", min_value)
		}
	}

	return mins, maxs, nil
}
