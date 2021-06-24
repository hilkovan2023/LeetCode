package main

import "fmt"

func count(child []int, cakes []int) int {
	res := 0
	i := len(child) - 1
	j := len(cakes) - 1
	for i >= 0 && j >= 0 {
		if cakes[j] >= child[i] {
			res++
			j--
			i--
		} else {
			i--
		}
	}
	return res
}

func main() {
	children := []int{1, 2, 7, 10, 11}
	cakesize := []int{1, 3, 5, 9}
	res := count(children, cakesize)
	fmt.Println(res)
}
