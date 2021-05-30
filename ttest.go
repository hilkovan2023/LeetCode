package main

import (
	"fmt"
	"runtime"
	"time"
)

func query() int {
	ch := make(chan int)
	for i := 0; i < 2; i++ {
		go func() { ch <- i }()
	}
	t := <-ch
	fmt.Println(t)
	return t
}

func main() {
	for i := 0; i < 4; i++ {
		query()
		time.Sleep(1 * time.Second)
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}
