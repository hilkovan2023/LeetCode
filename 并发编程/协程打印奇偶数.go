package main

import (
	"fmt"
	"sync"
	"time"
)

func one(in chan int) {
	for i := range in {
		fmt.Println(i + 1)
		time.Sleep(2000 * time.Millisecond)
	}
	fmt.Println("one ok!")
	defer wg.Done()
}

func two(out chan int) {
	for i := 0; i < 10; {
		fmt.Println(i)
		out <- i
		time.Sleep(500 * time.Millisecond)
		i += 2
	}
	close(out)
	fmt.Println(10)
	fmt.Println("two ok!")
	defer wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	ch1 := make(chan int)
	wg.Add(2)
	go one(ch1)
	go two(ch1)
	wg.Wait()
}
