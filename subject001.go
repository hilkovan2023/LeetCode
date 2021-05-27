package main

import "fmt"

func twoSum(nums []int, target int) []int {
	start := 0
	end := 0
	var k1 int
	var k2 int
	length := len(nums)
	for start = 0; start < length-1; start++{
		for end = start+1; end <length; end++{
			//fmt.Println(start, end)
			if nums[start] + nums[end] == target {
				k1 = start
				k2 = end
				break
			}
		}
	}
	//fmt.Println(start, end)
	return []int{k1, k2}
}

func main() {
	input := []int{1,2,3,4,5}
	target := 9
	output := twoSum(input, target)
	fmt.Println(output)
}
