package main

import "fmt"

func findnum(num []int, target int) int {
	if num == nil {
		return -1
	}
	if target < num[0] || target > num[len(num)-1] {
		return -1
	}
	left := 0
	right := len(num)
	min := left + (right-left)/2
	for {
		if num[min] == target {
			return min
		}
		if num[min] < target {
			left = min + 1
		}
		if num[min] > target {
			right = min
		}
		if left >= right {
			return -1
		}
		min = (left + right) / 2
	}
}

func main() {
	// 有序，无重复
	num := []int{-2, -1, 0, 1, 3, 9, 11, 19}
	res := findnum(num, -2)
	fmt.Println(res)
}
