package main

import "fmt"

var res = map[int]int64{}

func fib(n int) int64 {
	if n == 0 || n == 1 {
		res[n] = int64(n)
		return res[n]
	}

	if _, ok := res[n-1]; !ok {
		res[n-1] = fib(n - 1)
	}

	if _, ok := res[n-2]; !ok {
		res[n-2] = fib(n - 2)
	}
	res[n] = res[n-1] + res[n-2]
	return res[n]
}

func main() {
	N := 20
	fib(N)
	fmt.Println(fib(N))
	fmt.Println(res)
}
