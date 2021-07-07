package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	return fillhalf(nums, sum/2, val)
}

func fillhalf(nums []int, cap int, val []int) bool {
	fmt.Println("nums:", nums)
	fmt.Println("cap: ", cap)
	if cap < 0 {
		fmt.Println("cap<0", false)
		return false
	}
	if cap == 0 {
		fmt.Println("cap==0", true)
		fmt.Println("val", val)
		return true
	}
	if len(nums) == 0 {
		fmt.Println("len==0", false)
		return false
	}
	t1 := fillhalf(nums[1:], cap, val)

	val = append(val, nums[0])
	t2 := fillhalf(nums[1:], cap-nums[0], val)
	val = val[0 : len(val)-1]

	//fmt.Println("t2:", t2)
	fmt.Println("tt:", t1 && t2)
	return t1 || t2
}

var val = []int{}

func main() {
	nums := []int{2, 12, 4, 6, 3, 3}
	res := canPartition(nums)
	fmt.Println(res)
}
