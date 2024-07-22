package main

import (
	"fmt"
	"gift/database"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func init() {

}

//分布式锁

func TryLock(rc *redis.Client, name string, life time.Duration) bool {
	cmd := rc.SetNX(name, 1, life) //setNX
	if cmd.Err() != nil {
		return false
	}
	return cmd.Val()
}
func ReleaseLock(rc *redis.Client, name string) {
	rc.Del(name)
}

func main() {
	rc := database.GetRedisClient()
	const P = 100
	const lockName = "dpp"
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done()
			if TryLock(rc, lockName, 3*time.Millisecond) {
				fmt.Printf("上锁成功\n")
			}
		}()
	}
	wg.Wait()
	ReleaseLock(rc, lockName)
}
