package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxSubArray(nums []int, l int, sum []int) int {
	if l == 0 {
		return 0
	}
	if l == 1 {
		return Max(0, nums[0])
	}

	sum[l-2] = MaxSubArray(nums, l-1, sum)
	max := Max(sum[l-2]+nums[l-1], 0)
	sum[l-1] = max
	fmt.Println("sum", sum)
	fmt.Println("max", max, l)
	return max
}

func main() {
	//nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -2, 4}
	fmt.Println(nums)
	sum := make([]int, len(nums))
	res := MaxSubArray(nums, len(nums), sum)
	fmt.Println(res)
	fmt.Println(sum)
}
