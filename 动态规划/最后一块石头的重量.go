package main

import "fmt"

func lastStoneWeightII(weights []int) int {
	sum := 0
	for _, v := range weights {
		sum += v
	}
	minRestWeight = sum / 2
	res := maxWeightInBag(weights, sum/2)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 一堆石头weights，袋子容量bagWeight，袋子最满放多重的石头
func maxWeightInBag(weights []int, bagWeight int) int {

	if bagWeight == 0 {
		minRestWeight = bagWeight
		fmt.Println("loads:", bagWeight, loads)
		return 0
	}
	if len(weights) == 0 {
		if bagWeight < minRestWeight {
			minRestWeight = bagWeight
			fmt.Println("weights:", bagWeight, loads)
		}
		return 0
	}

	if weights[0] > bagWeight {
		return maxWeightInBag(weights[1:], bagWeight)
	} else {
		loads = append(loads, weights[0])
		m1 := maxWeightInBag(weights[1:], bagWeight-weights[0]) + weights[0]
		loads = loads[:len(loads)-1]

		m2 := maxWeightInBag(weights[1:], bagWeight)

		return max(m1, m2)
	}
}

var loads = []int{}
var minRestWeight = 0

func main() {
	weights := []int{2, 5, 2, 1}
	res := lastStoneWeightII(weights)
	fmt.Println(res)
}
