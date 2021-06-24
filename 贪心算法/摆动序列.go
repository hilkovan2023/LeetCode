package main

import "fmt"

func maxswinglen(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	diff := nums[1] - nums[0]
	fast := 2
	slow := 1
	res := 2
	for fast < len(nums) {
		if diff*(nums[fast]-nums[slow]) >= 0 {
			fast++
			slow++
		} else {
			diff = nums[fast] - nums[slow]
			res++
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, 3, 5, 7, 9, 11, 3, 2, 11, 10, 15, 14, 13}
	res := maxswinglen(nums)
	fmt.Println(res)
}
