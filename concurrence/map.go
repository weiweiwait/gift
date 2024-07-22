package main

import (
	"sync"
	"time"
)

var mp = make(map[int]int, 1000)
var smp = sync.Map{} //本身支持并发读写的map
func readMap() {
	for i := 0; i < 1000; i++ {
		//lock.RLock()
		//_ = mp[10]
		//lock.RUnlock()
		smp.Load(10)
	}
}
func writeMap() {
	for i := 0; i < 1000; i++ {
		//lock.Lock()
		//mp[5] = 5
		//lock.Unlock()
		smp.Store(5, 5)
	}

}

// 使用sync.Map计数会存在问题(脏写)
func mapInc(mp *sync.Map, key int) {
	if oldValue, exists := mp.Load(key); exists {
		mp.Store(key, oldValue.(int)+1)
	} else {
		mp.Store(key, 1)
	}
}
func main() {
	go readMap()
	go writeMap()
	time.Sleep(1 * time.Second)
}
