package main

import (
	"fmt"
	"math/rand"
	"time"
)

func product(out chan<- int, index int) {
	for {
		// 判断缓冲区是否满
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("生产者%dth,生产:%d\n", index, num)
		// 访问公共区结束，并且打印结束，解锁

		// 唤醒阻塞在条件变量上的 消费者
		time.Sleep(time.Millisecond * 300)
	}
}

func consumer(in <-chan int, index int) {
	for {
		// 先加锁

		// 判断 缓冲区是否为空

		num := <-in
		fmt.Printf("----消费者%dth,消费:%d\n", index, num)
		// 访问公共区结束后，解锁

		// 唤醒 阻塞在条件变量上的 生产者

		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	ch := make(chan int, 5)
	quit := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	// 指定条件变量 使用的锁

	for i := 0; i < 5; i++ {
		go product(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go consumer(ch, i+1)
	}
	<-quit
}
