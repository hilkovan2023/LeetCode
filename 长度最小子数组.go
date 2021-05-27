package main

import "fmt"

func minlist(num []int, target int) int {
	// 左闭右开
	minslow := 0
	minfast := len(num)
	slow := 0
	sum := 0
	for fast := 1; fast <= len(num); fast++ {
		sum += num[fast-1]
		if sum > target {
			for sum > target {
				sum -= num[slow]
				slow++
			}
			fmt.Println(slow-1, fast)
			if fast-slow+1 < minfast-minslow {
				minfast = fast
				minslow = slow - 1
			}
		}
	}
	return minfast - minslow
}

func main() {
	num := []int{1, 2, 3, 2, 5, 3, 2, 2, 3}
	res := minlist(num, 10)
	fmt.Println(res)
}
