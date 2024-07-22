package main

import (
	"fmt"
	"sync"
)

// slice支持并发操作，但是并不保证结果准确性
func main() {
	const LEN = 10
	arr := make([]int, LEN)
	const P = 2
	const LOOP = 10000
	wg := sync.WaitGroup{}
	wg.Add(P)
	//for i := 0; i < P; i++ {
	go func() {
		defer wg.Done()
		for j := 0; j < LOOP; j++ {
			for index := 0; index < LEN; index++ {
				if index <= LEN/2 {
					arr[index]++
				}
			}
		}
	}()
	//}
	go func() {
		defer wg.Done()
		for j := 0; j < LOOP; j++ {
			for index := 0; index < LEN; index++ {
				if index > LEN/2 {
					arr[index]++
				}
			}
		}
	}()
	wg.Wait()
	sum := 0
	for _, ele := range arr {
		sum += ele
	}
	fmt.Println(sum)
}
