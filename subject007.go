package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32{
		return 0
	}
	sign := true
	output := 0
	if x < 0 {
		sign = false
		x = -x
	}
	yu := 0
	for x >= 1 {
		yu = x%10
		x /= 10
		output *= 10
		output += yu
	}
	if output > math.MaxInt32 {
		return 0
	}
	if sign == false {
		output = -output
	}
	return output
}

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}