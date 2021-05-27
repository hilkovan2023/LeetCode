package main

import "fmt"


func addOne(s []int) {
	s[0] = 4  // 可以改变原切片值
	s = append(s, 1)  // 扩容后分配了新的地址，原切片将不再受影响
	s[0] = 8
	fmt.Println(s)
}

type nnn struct {
	val int
	left *nnn
	right *nnn
}



func main() {
	p := new(nnn)
	q := nnn{val:100}
	fmt.Println(p)
	fmt.Println(q)
	var s1 = []int{2}   // 初始化一个切片
	fmt.Println(s1)
	addOne(s1)          // 调用函数添加一个切片
	fmt.Println(s1)     // 输出一个值 [4]
}
