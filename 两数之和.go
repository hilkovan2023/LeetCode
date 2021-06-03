package main

import "fmt"

func find_two_num(num []int, target int) (int, int) {
	m := make(map[int]int)
	for i, v := range num {
		if _, ok := m[target-v]; ok {
			return m[target-v], i
		}
		m[v] = i
	}
	return -1, -1
}

func main() {
	num := []int{1, 2, 3, 4, 5, 22, 12, 14, 22}
	idx1, idx2 := find_two_num(num, 44)
	fmt.Println(idx1, idx2)
}
