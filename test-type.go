package main

import (
	"fmt"
)

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	// 开启goroutine将0~100的数发送到ch1中
//	go func() {
//		for i := 0; i < 10; i++ {
//			time.Sleep(time.Second)
//			ch1 <- i
//			fmt.Println("ch1 <- :", i)
//		}
//		close(ch1)
//	}()
//	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	go func() {
//		for {
//			i, ok := <-ch1 // 通道关闭后再取值ok=false
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//			fmt.Println("ch2 <- :", i*i)
//		}
//		close(ch2)
//	}()
//	// 在主goroutine中从ch2中接收值打印
//	for i := range ch2 { // 通道关闭后会退出for range循环
//		fmt.Println("range ch2", i)
//
//	}
//}

func xiao(a []int) {
	fmt.Printf("%p, %p, %p", len(a), cap(a), a)
}
func main() {
	//a := new([]int)
	//fmt.Println(a)  //输出&[]，ａ本身是一个地址

	a := make([]int, 5)
	a = append(a, 1)
	fmt.Printf("%p, %p, %p", len(a), cap(a), a)
	xiao(a)
	fmt.Printf("%p, %p, %p", len(a), cap(a), a)
	fmt.Println()
	b := make([]int, 1)
	fmt.Println(b) //输出[0]，ｂ本身是一个slice对象，其内容默认为０
	c := make(map[int]int)
	c[3] = 3
	c[4] = 4
	fmt.Println(c)
	fmt.Printf("%v   %v   %v ", a, b, c)

	m := new(int)

	fmt.Println(m)

	var l int
	fmt.Println(&l)
	fmt.Println(&l)
	l = 10
	fmt.Printf("%v  %p", l, l)
}
