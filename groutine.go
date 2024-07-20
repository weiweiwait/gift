package main

import (
	"fmt"
	"time"
)

func parent() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func child() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() { //runtime来调度协程，所以main结束了就全结束了
	//main是一个特殊的协程
	go parent()
	go child()
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("main")
}
