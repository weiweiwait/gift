package database

import (
	"log"
	"strconv"
)

const (
	prefix = "gift_count_"
)

//从Mysql中读出所有奖品的初始库存,存入Redis,如果同时有很多用户来参与抽奖，不能直接在mysql里面减库存,mysql扛不住折磨搞得并发量

func InitGiftInventory() {

}

//获取所有奖品剩余的库存量

func GetAllGiftInventory() []*Gift {
	client := GetRedisClient()
	keys, err := client.Keys(prefix + "*").Result()
	if err != nil {
		//有待改变
		log.Fatalf("iterate all keys by prefix %s failed: %s\n", prefix, err)
		return nil
	}
	gifts := make([]*Gift, 0, len(keys))
	for _, key := range keys {
		if id, err := strconv.Atoi(key[len(prefix):]); err == nil {
			count, err := client.Get(key).Int()
			if err == nil {
				gifts = append(gifts, &Gift{Id: id, Count: count})
			} else {
				log.Fatalf("invalid gift inventory %s", client.Get(key).String())
			}
		} else {
			log.Fatalf("invalid redis key %s", key)
		}
	}
	return gifts
}
