package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup完成子携程结束，直接main终止，不用手动
var (
	wg = sync.WaitGroup{}
)

func init() {
	wg.Add(2)
}
func parent() {
	defer wg.Done() //所有执行完后执行
	go child()
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%d\n", i)
		//time.Sleep(500 * time.Millisecond)
	}
}
func child() {
	defer wg.Done()
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() { //runtime来调度协程，所以main结束了就全结束了
	//main是一个特殊的协程
	//go parent()
	go func() {
		defer wg.Done() //所有执行完后执行
		go child()
		for i := 'a'; i < 'z'; i++ {
			fmt.Printf("%d\n", i)
			//time.Sleep(500 * time.Millisecond)
		}
	}()
	go child()
	fmt.Println("main")
	//time.Sleep(1500 * time.Millisecond)
	wg.Wait() //等wg减为0
}
