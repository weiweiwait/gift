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
func main() {
	go parent()
	go child()
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("main")
}
