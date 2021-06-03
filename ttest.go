package main

import (
	"fmt"
	"time"
)

func read1(int_ch chan int) {
	for {
		rx := <-int_ch
		fmt.Println("read1:", rx)
	}
}

func read2(int_ch chan int) {
	for {
		rx := <-int_ch
		fmt.Println("read2:", rx)
	}
}

func read3(int_ch chan int) {
	for {
		rx := <-int_ch
		fmt.Println("read3:", rx)
	}
}

// 判断管道有没有存满
func main() {
	int_ch := make(chan int)

	go read1(int_ch)

	go read2(int_ch)

	go read3(int_ch)

	for i := 0; i >= 0; i++ {
		int_ch <- i
		time.Sleep(1 * time.Second)
	}
}
