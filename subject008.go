package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	if str == "" {return 0}

	index := 0
	for  i:= 0; i < len(str); i++{
		if str[i] != ' '{
			index = i
			break
		}
	}

	sign := true
	if str[index] == '-' {
		sign = false
		index++
	} else if str[index] == '+' {
		sign = true
		index++
	}

	if index == len(str) {
		return 0
	}

	if str[index] > 57 || str[index] < 48 {
		return 0
	}

	var output int64 = 0
	for j := index; j < len(str); j++ {
		if str[j] <= 57 && str[j] >= 48 {
			output *= 10
			output += int64((str[j]-48))
			if output > math.MaxInt32+1 {
				break
			}
		} else {
			break
		}
	}

	if sign {
		if output > math.MaxInt32 {
			output = math.MaxInt32
		}
	} else {
		output = -output
		if output < math.MinInt32 {
			output = math.MinInt32
		}
	}
	return int(output)
}

func main()  {
	str := "9999991936854775808"
	fmt.Println(myAtoi(str))
}