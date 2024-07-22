package database

import (
	"fmt"
	"log"
	"strconv"
)

const (
	prefix = "gift_count_"
)

//从Mysql中读出所有奖品的初始库存,存入Redis,如果同时有很多用户来参与抽奖，不能直接在mysql里面减库存,mysql扛不住折磨搞得并发量

func InitGiftInventory() {
	giftCh := make(chan Gift, 100)
	//go GetAllGiftsV1()
	client := GetRedisClient()
	for {
		gift, ok := <-giftCh
		if !ok {
			//channel已经消费完了
			break
		}
		if gift.Count <= 0 {
			continue //没有库存的商品不参与抽奖
		}
		err := client.Set(prefix+strconv.Itoa(gift.Id), gift.Count, 0).Err()
		if err != nil {
			log.Fatalf("set gift %d:%s count to %d failed: %s", gift.Id, gift.Name, gift.Count, err)
		}
	}
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

//奖品对应的库存-1

func ReduceInventory(GiftId int) error {
	client := GetRedisClient()
	key := prefix + strconv.Itoa(GiftId)
	n, err := client.Decr(key).Result()
	if err != nil {
		log.Fatalf("decr key %s failed %s\n", key, err)
		return err
	} else {
		if n < 0 {
			log.Printf("%d一无库存，减少失败", GiftId)
			return fmt.Errorf("%d一无库存，减少失败", GiftId)
		}
		return nil
	}
}
