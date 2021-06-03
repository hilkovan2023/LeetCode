package main

import "fmt"

func intersection(list1 []int, list2 []int) []int {
	map1 := make(map[int]struct{})
	map2 := make(map[int]struct{})
	for _, v1 := range list1 {
		map1[v1] = struct{}{}
	}
	for _, v2 := range list2 {
		if _, ok := map1[v2]; ok {
			map2[v2] = struct{}{}
		}
	}
	res := []int{}
	for k, _ := range map2 {
		res = append(res, k)
	}
	return res
}

func main() {
	list1 := []int{1, 2, 2, 1}
	list2 := []int{2, 2}
	res := intersection(list1, list2)
	fmt.Println(res)
}
