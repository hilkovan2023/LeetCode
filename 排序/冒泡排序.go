package main

import "fmt"

func sort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{1, 3, 6, 2, 5, 4, 8, 2, 6}
	sort(nums)
	fmt.Println(nums)
}
