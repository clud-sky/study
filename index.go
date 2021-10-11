package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func zhoulin(zl chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(zl <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-zl
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}

func testMain() {
	wg.Add(1)
	go zhoulin(jobChan)

	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.result)
	}
	wg.Wait()
}

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	//ch := make(chan int,1)
	//for i:=0;i<10;i++{
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x)
	//	case ch<- i:
	//	}
	//}

	f1 := float32(1.23456)
	fmt.Printf("%T\n", f1)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%2f\n", math.Pi)

	f2 := "12345978642"
	ass := strings.Contains(f2, "0")
	fmt.Println(ass)

}

func testWg() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func hello(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func testGomaxprocs() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(&stu)
	}

	fmt.Println(m)
}

func testDeferCall() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("no")
		}
	}()

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发一场")
}

var c1 = make(chan int)
var c2 = make(chan int)

func testChan() {
	go func() {

		fmt.Println("1111")
		c2 <- 1
		c1 <- <-c2

	}()

	go func() {
		c1 <- <-c2
		fmt.Println("2222")
	}()
	<-c1
}

func testDefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("no")
		}
	}()

	defer func() {
		panic("1111")
	}()

	panic("2222")
}
