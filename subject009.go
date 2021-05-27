package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	yy := 0
	xx := 0
	if x < 0 {
		xx = -x
	} else {
		xx = x
	}
	for xx >= 1 {
		yy *= 10
		yy += xx%10
		xx /= 10
		//fmt.Println("xx",xx)
		//fmt.Println("yy",yy)
	}

	if yy == x {
		return true
	} else {
		return false
	}
}

func main()  {
	x := -123434321
	fmt.Println(isPalindrome(x))
}