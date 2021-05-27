package main

import "fmt"

func luoxuan(n int) [][]int {

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left := 0
	right := n - 1
	top := 0
	bot := n - 1
	y := 0
	x := 0
	i := 1
	for i <= n*n {

		for x = left; x <= right; x++ {
			res[top][x] = i
			i++
		}
		top++
		fmt.Println(res)

		for y = top; y <= bot; y++ {
			res[y][right] = i
			i++
		}
		right--
		fmt.Println(res)

		for x = right; x >= left; x-- {
			res[bot][x] = i
			i++
		}
		bot--
		fmt.Println(res)

		for y = bot; y >= top; y-- {
			res[y][left] = i
			i++
		}
		left++
		fmt.Println(res)
	}
	return res
}

func main() {
	res := luoxuan(9)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
