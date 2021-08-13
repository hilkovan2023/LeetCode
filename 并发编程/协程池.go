package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func newjob(jobch chan *Job) {
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{id, r_num}
		jobch <- job
	}
}

func printres(resch chan *Result) {
	for res := range resch {
		fmt.Printf("job id:%v randnum:%v result:%d\n", res.job.Id, res.job.RandNum, res.sum)
	}
}

func worker(jobch chan *Job, resch chan *Result) {
	for job := range jobch {
		// 随机数接过来
		r_num := job.RandNum
		// 随机数每一位相加
		// 定义返回值
		var sum int
		for r_num != 0 {
			tmp := r_num % 10
			sum += tmp
			r_num /= 10
		}
		// 想要的结果是Result
		r := &Result{
			job: job,
			sum: sum,
		}
		//运算结果扔到管道
		resch <- r
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go worker(jobChan, resultChan)
	}
}

// goroutin()不能放在最后，防止main()先退出
// channel用于goroutin间通信
func main() {

	// job管道
	jobChan := make(chan *Job, 128)
	// 结果管道
	resultChan := make(chan *Result, 128)

	// 循环创建job，输入到管道
	go newjob(jobChan)
	// 创建工作池
	createPool(64, jobChan, resultChan)
	// 开个打印的协程
	printres(resultChan)
}
