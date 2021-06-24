package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	res := []int{}
	for i := 0; i < k; i++ {
		if len(queue) == 0 || nums[i] <= queue[len(queue)-1] {
			queue = append(queue, nums[i])
		} else if nums[i] > queue[0] {
			queue = queue[0:0]
			queue = append(queue, nums[i])
		} else {
			for nums[i] > queue[len(queue)-1] {
				queue = queue[:len(queue)-1]
			}
			queue = append(queue, nums[i])
		}
		fmt.Println("queue", queue)
	}
	res = append(res, queue[0])

	for j := k; j < len(nums); j++ {
		if nums[j-k] == queue[0] {
			queue = queue[1:]
		}
		if nums[j] > queue[0] {
			queue = queue[0:0]
			queue = append(queue, nums[j])
		} else if nums[j] <= queue[len(queue)-1] {
			queue = append(queue, nums[j])
		} else {
			for nums[j] > queue[len(queue)-1] {
				queue = queue[:len(queue)-1]
			}
			queue = append(queue, nums[j])
		}

		fmt.Println("queue", queue)
		res = append(res, queue[0])
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 4
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
