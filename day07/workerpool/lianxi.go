package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
   使用goroutine和channel实现一个计算int64随机数各位数和的程序。
       开启一个goroutine循环生成int64类型的随机数，发送到jobChan
       开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
       主goroutine从resultChan取出结果并打印到终端输出
   为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。
*/

type Job struct {
	value int64
}

type Result struct {
	job *Job
	id  int
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func zhoulin(zl chan<- *Job) {
	defer wg.Done()
	// 选好生成int64类型发送到jobChan
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(id int, zl <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	for {
		job := <-zl
		var sum int64
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		newResult := &Result{
			id:  id,
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go zhoulin(jobChan)
	// 开启24个goroutine执行取值
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(i, jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("worker:%d value:%d sum:%d\n", result.id, result.job.value, result.sum)
	}

	wg.Wait()
}
