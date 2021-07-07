package main

import "fmt"

// 成功nums中找到满足target的组合数量
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	left := (target + sum) / 2
	fmt.Println("left:", left)
	res := findWays(nums, left)
	return res
}

func findWays(nums []int, target int) int {
	if target == 0 {
		fmt.Println("sign:", sign)
		return 1
	}
	if target < 0 {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	sign = append(sign, 1)
	r1 := findWays(nums[1:], target-nums[0])
	sign = sign[:len(sign)-1]
	r2 := findWays(nums[1:], target)
	return r1 + r2
}

var sign = []int{} // 符号组合

func main() {
	nums := []int{1, 1, 1, 1, 1}
	res := findTargetSumWays(nums, 3)
	fmt.Println(res)
}
