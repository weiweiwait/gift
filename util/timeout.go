package util

import (
	"context"
	"fmt"
	"time"
)

const (
	WorkUseTime = 500 * time.Millisecond
	Timeout     = 100 * time.Millisecond
)

//模拟一个耗时较长的任务

func LongTimeWork() int {
	time.Sleep(WorkUseTime)
	return 888
}

//模拟一个接口处理函数

func Handle1() int {
	deadline := make(chan struct{}, 1)
	workDone := make(chan int, 1)
	go func() { //把要超时控制的函数放一个携程里面
		n := LongTimeWork()
		workDone <- n
	}()
	go func() { //把要超时控制的函数放一个携程里面
		time.Sleep(Timeout)
		deadline <- struct{}{}
		//close(deadline)
	}()
	select {
	case n := <-workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <-deadline:
		fmt.Println("LongTimeWork timeout")
		return 0
	}

}
func Handle2() int {
	workDone := make(chan int, 1)
	go func() { //把要超时控制的函数放一个携程里面
		n := LongTimeWork()
		workDone <- n
	}()
	select {
	case n := <-workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <-time.After(Timeout):
		fmt.Println("LongTimeWork timeout")
		return 0
	}

}

func Handle3() int {
	workDone := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() { //把要超时控制的函数放一个携程里面
		n := LongTimeWork()
		workDone <- n
	}()
	go func() { //把要超时控制的函数放一个携程里面
		time.Sleep(Timeout)
		//close(deadline)
		cancel()
	}()
	select {
	case n := <-workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <-ctx.Done():
		fmt.Println("LongTimeWork timeout")
		return 0
	}
}
