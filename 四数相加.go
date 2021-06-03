package main

import "fmt"

func fournumsum(num1 []int, num2 []int, num3 []int, num4 []int) int {
	res := 0
	m12 := make(map[int]int)
	for _, v1 := range num1 {
		for _, v2 := range num2 {
			if _, ok := m12[v1+v2]; ok {
				m12[v1+v2]++
			} else {
				m12[v1+v2] = 1
			}
		}
	}
	fmt.Println(m12)

	for _, v3 := range num3 {
		for _, v4 := range num4 {
			sum34 := -v3 - v4
			if _, ok := m12[sum34]; ok {
				res += m12[sum34]
			}
		}
	}
	return res
}

func main() {
	num1 := []int{1, 2}
	num2 := []int{-1, -2}
	num3 := []int{-1, 2}
	num4 := []int{0, 2}
	res := fournumsum(num1, num2, num3, num4)
	fmt.Println(res)
}
