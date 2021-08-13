package main

import "fmt"

func dailyTemperatures(temp []int) []int {
	res := make([]int, len(temp))
	stack := make([]int, len(temp)) // 栈
	stack[0] = 0                    // temp第一个元素的下标入栈
	top := 0                        // 栈顶
	for i := 1; i < len(temp); i++ {
		for temp[i] > temp[stack[top]] && top >= 0 {
			res[stack[top]] = i - stack[top]
			stack = stack[0 : len(stack)-1]
			top--
			if top < 0 {
				break
			}
		}
		top++
		stack[top] = i
	}

	return res
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73, 12, 22, 32, 15, 23, 44}
	//fmt.Println(len(temperatures), len(temperatures[0:0]))
	results := dailyTemperatures(temperatures)
	fmt.Println("res:", results)
}
