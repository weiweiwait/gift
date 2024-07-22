package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

// 读锁和写锁总是互斥的，读锁之间互相不排斥，写锁之间互相排斥
func main() {
	lock.RLock() //上写锁
	go func() {
		lock.Lock() //上写锁
		fmt.Println("上锁成功")
	}()
	time.Sleep(3 * time.Second)
}
