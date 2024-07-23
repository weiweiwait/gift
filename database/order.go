package database

import "log"

type Order struct {
	Id     int
	GiftId int
	UserId int
}

//写入一条订单记录

func CreateOrder(userid, giftid int) int {
	db := GetGiftDBConnection()
	order := Order{
		GiftId: giftid,
		UserId: userid,
	}
	if err := db.Create(&order).Error; err != nil {
		log.Printf("create order failed: %s", err)
		return 0
	} else {
		log.Printf("create order id is %d", order.Id)
		return order.Id
	}
}

//清除全部订单记录

func ClearOrders() error {
	db := GetGiftDBConnection()
	return db.Where("id>0").Delete(Order{}).Error
}
