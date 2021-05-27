package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	tmax := 0
	for start := 0; start < len(height)-1; start++ {
		for end := start+1; end < len(height); end++ {
			if height[start] < height[end] {
				tmax = height[start]*(end-start)
				//fmt.Println("tmax=", tmax)
			} else {
				tmax = height[end]*(end-start)
			}
			if tmax > max {
				//fmt.Println(start, end)
				max = tmax
			}
		}
	}
	return max
}

func main()  {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}
