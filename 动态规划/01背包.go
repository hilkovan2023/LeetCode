package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//
func wei_bag_problem(weight []int, value []int, bagWeight int) int {
	fmt.Println("len(weight): ", len(weight))
	if len(weight) == 0 {
		return 0
	}
	if bagWeight <= 0 {
		return 0
	}
	if v, ok := maxval[bagWeight]; ok {
		return v
	}
	v_load := 0
	if bagWeight-weight[0] >= 0 {
		v_load = wei_bag_problem(weight[1:], value[1:], bagWeight-weight[0]) + value[0]
	} else {
		v_load = 0
	}

	v_unload := wei_bag_problem(weight[1:], value[1:], bagWeight)
	m_val := max(v_load, v_unload)
	maxval[bagWeight] = m_val
	fmt.Println("m_val: ", m_val)
	return m_val
}

// map[w] = v   背包容量为w时，最大价值v
var maxval map[int]int

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagWeight := 100
	maxval = make(map[int]int)
	res := wei_bag_problem(weight, value, bagWeight)
	fmt.Println(res)
	fmt.Println("maxval:", maxval)
}
