package main

import "fmt"

func rmele1(num []int, target int) int {
	if num == nil {
		return 0
	}
	last := len(num) - 1
	for i := 0; i <= last; i++ {
		for num[last] == target {
			last--
		}
		if num[i] == target {
			num[i], num[last] = num[last], num[i]
			last--
		}
	}
	return last + 1
}

func rmele2(num []int, target int) int {
	if num == nil {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(num); fast++ {
		if num[fast] != target {
			num[slow] = num[fast]
			slow++
		}

	}
	return slow
}

func main() {
	num := []int{0, 11, 11, 32, 23, 1, 31, 5, 30, 3, 23, 9}
	fmt.Println(" num ", num, len(num))
	res1 := rmele1(num, 11)
	fmt.Println("func1", num, res1)
	res2 := rmele2(num, 11)
	fmt.Println("func2", num, res2)
}
