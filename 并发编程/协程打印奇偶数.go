package main

import (
	"fmt"
	"sync"
)

type Printer struct {
	wg  *sync.WaitGroup
	ch1 chan int
	ch2 chan int
}

func NewPrinter(wg *sync.WaitGroup) Printer {
	return Printer{
		wg:  wg,
		ch1: make(chan int),
		ch2: make(chan int),
	}
}

func (p Printer) Start(start int) {
	p.ch1 <- start
}

func (p Printer) One() {
	for i := range p.ch1 {
		fmt.Println(i)
		p.ch2 <- i + 1
	}
	close(p.ch2)
	fmt.Println("one ok!")
	defer p.wg.Done()
}

func (p Printer) Two() {
	for i := range p.ch2 {
		if i > 16 {
			close(p.ch1)
			return
		}
		fmt.Println(i)
		p.ch1 <- i + 1
	}
	fmt.Println("two ok!")
	defer p.wg.Done()
}

func main() {
	var wg sync.WaitGroup
	printer := NewPrinter(&wg)
	wg.Add(2)
	go printer.One()
	go printer.Two()

	printer.Start(3)
	wg.Wait()
}
