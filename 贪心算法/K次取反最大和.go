package main

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func main() {
	nums := []int{2, -3, -1, 5, -4, 11, 13, -21, 63, 19}
	//sort.Slice(nums, func(i, j int) bool {
	//	return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	//})
	sum := largestSumAfterKNegations(nums, 9)
	fmt.Println(sum)
}
