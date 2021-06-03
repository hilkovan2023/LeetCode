package main

import "fmt"

func squsum(num int) int {
	res := 0
	for num > 0 {
		yu := num % 10
		num = num / 10
		res += yu * yu
	}
	return res
}

func ishappynum(num int) bool {
	if num <= 0 {
		return false
	}

	m := make(map[int]struct{})
	for num != 1 {
		fmt.Println(num)
		if _, ok := m[num]; ok {
			return false
		}
		m[num] = struct{}{}
		num = squsum(num)
	}
	return true
}

func main() {
	num := 82
	res := ishappynum(num)
	fmt.Println(res)
}
