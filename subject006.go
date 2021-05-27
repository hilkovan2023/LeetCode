package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) == 0 || len(s) == 1{
		return s
	}
	if numRows == 1 {
		return s
	}

	temp := make([]uint8, len(s))
	index := 0
	for m := 0; m*(2*numRows-2) < len(s); m++ {
		temp[index] = s[(2*numRows - 2) * m]
		index++
	}
	fmt.Println("start", string(temp))
	for n := 1; n < numRows-1; n++ {
		for m := 0; m*(2*numRows-2) < len(s); m++ {
			if index < len(s) {
				if (2*numRows-2)*m + n < len(s) {
					temp[index] = s[(2*numRows-2)*m + n]
					index++
				}
				if (m+1)*(2*numRows-2) - n < len(s) {
					temp[index] = s[(m+1)*(2*numRows-2) - n]
					index++
				}
				//fmt.Println((2*numRows-2)*m + n, (m+1)*(2*numRows-2) - n)
				//fmt.Println(string(temp))
			}
		}
	}
	fmt.Println("middle", string(temp))
	for m := 0; m*(2*numRows-2) <= len(s); m++ {
		if (2*numRows-2)*m + numRows - 1 < len(s) {
			temp[index] = s[(2*numRows-2)*m + numRows - 1]
			index++
		}
	}
	fmt.Println("end", string(temp))
	return string(temp)
}

func main() {
	s := "ag"
	fmt.Println(convert(s, 1))
}