package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool 协程池结构
type WorkerPool struct {
	workerNum int            // 工作协程数量
	jobQueue  chan JOB       // 任务队列
	wg        sync.WaitGroup // 等待组，用于等待所有任务完成
}

// Job 任务接口
type JOB interface {
	Do()
}

// 示例任务
type SimpleJob struct {
	id int
}

func (j SimpleJob) Do() {
	fmt.Printf("Job %d started\n", j.id)
	time.Sleep(time.Second) // 模拟耗时工作
	fmt.Printf("Job %d finished\n", j.id)
}

// NewWorkerPool 创建新的协程池
func NewWorkerPool(workerNum int, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		workerNum: workerNum,
		jobQueue:  make(chan JOB, queueSize),
	}

	// 启动工作协程
	pool.start()
	return pool
}

// start 启动工作协程
func (p *WorkerPool) start() {
	for i := 0; i < p.workerNum; i++ {
		go func(workerID int) {
			for job := range p.jobQueue {
				job.Do()
				p.wg.Done()
			}
		}(i)
	}
}

// Submit 提交任务
func (p *WorkerPool) Submit(job JOB) {
	p.wg.Add(1)
	p.jobQueue <- job
}

// Wait 等待所有任务完成
func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

// Shutdown 关闭协程池
func (p *WorkerPool) Shutdown() {
	close(p.jobQueue)
}

func main() {
	// 创建一个有3个工作协程、任务队列容量为10的协程池
	pool := NewWorkerPool(3, 10)

	// 提交10个任务
	for i := 1; i <= 10; i++ {
		job := SimpleJob{id: i}
		pool.Submit(job)
	}

	// 等待所有任务完成
	pool.Wait()

	// 关闭协程池
	pool.Shutdown()

	fmt.Println("All jobs completed")
}
