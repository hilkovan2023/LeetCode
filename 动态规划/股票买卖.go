package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 第n天，可以获得的最大收益, has表示是否持有股票
func maxProfit(prices []int, n int, has bool) int {
	if n == 0 && has == true {
		return -prices[0]
	}
	if n == 0 && has == false {
		return 0
	}
	if len(prices) <= 1 {
		return 0
	}

	// 持有
	if has == true {
		p1 := maxProfit(prices, n-1, false) - prices[n]
		p2 := maxProfit(prices, n-1, true)

		res := max(p1, p2)
		return res
	} else {
		// 不持有
		p1 := maxProfit(prices, n-1, true) + prices[n]
		p2 := maxProfit(prices, n-1, false)

		res := max(p1, p2)
		return res
	}
}

var prices = []int{1, 5, 3, 7}

func main() {
	// 二叉树结构，根据参数has分叉，n表示层数
	res := maxProfit(prices, len(prices)-1, false)
	fmt.Println(res)
}
