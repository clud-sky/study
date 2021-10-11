package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	f1(ctx)
FORLOOP:
	for {
		fmt.Println("刘洋君")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}

func f1(ctx context.Context) {
	defer wg.Done()
loop:
	for {
		fmt.Println("大傻逼")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break loop
		default:

		}
	}
}

func first() {
	ctx, cencel := context.WithCancel(context.Background())

	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 3)
	cencel()
	wg.Wait()
}

func deadline() {
	d := time.Now().Add(2000 * time.Millisecond)
	ctx, cencel := context.WithDeadline(context.Background(), d)
	defer cencel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("刘洋君")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	deadline()
}
