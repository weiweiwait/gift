package database

import (
	"gift/util"
	"gorm.io/gorm"
	"log"
)

const EMPTY_GIFT = 1 //空奖品("谢谢参与")的ID
type Gift struct {
	Id      int    `gorm:"column:id;primaryKey"`
	Name    string `gorm:"column:name"`
	Price   int    `gorm:"column:price"`
	Picture string `gorm:"column:picture"`
	Count   int    `gorm:"column:count"`
}

func (g Gift) TableName() string {
	return "inventory"
}

var (
	_all_gift_field = util.GetGormFields(Gift{})
)

// 把gift表里面数据全读出来，当数量不多时可以直接select *

func GetAllGiftsV1() []*Gift {
	db := GetGiftDBConnection()
	var gifts []*Gift
	err := db.Select(Gift{}).Find(&gifts).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			//util.LogRus.Errorf("read table %s failed: %s", Gift{}.TableName(), err)
			log.Fatalf("read table %s failed: %s", Gift{}.TableName(), err)
		}
	}
	return gifts
}

//千万级以上大表便利方案

func GetAllGiftsV2(ch chan<- Gift) {
	db := GetGiftDBConnection()
	const PAGE_SIZE = 500
	maxid := 0
	for {
		var gifts []Gift
		err := db.Select(Gift{}).Where("id>?", maxid).Limit(PAGE_SIZE).Find(&gifts).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Fatalf("read table %s failed %s", Gift{}.TableName(), err)
			}
			break
		}
		if len(gifts) == 0 {
			break
		}
		for _, gifts := range gifts {
			ch <- gifts
			// 更新maxID为当前页最大的ID，以便下一页查询
			if gifts.Id > maxid {
				maxid = gifts.Id
			}
		}
	}
	close(ch)
}
