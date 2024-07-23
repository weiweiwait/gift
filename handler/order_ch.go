package handler

import (
	"gift/database"
	"log"
)

var (
	orderch = make(chan database.Order, 10000)
	stopCh  = make(chan struct{}, 1)
)

func InitChannel() {
	go func() {
		<-stopCh
		close(orderch)
	}()
}

//把订单放入channel

func PutOrder(UserId, GiftId int) {
	order := database.Order{
		UserId: UserId,
		GiftId: GiftId,
	}
	orderch <- order
}

//读订单取出来进行入库操作

func TakeOrder() {
	for {
		order, ok := <-orderch
		if !ok {
			log.Println("order channel is empty and closed")
			break
		}
		database.CreateOrder(order.UserId, order.GiftId)
	}
}

// 目的是关闭orderCh , 该函数可以反复调用

func CloseChannel() {
	stopCh <- struct{}{}
	select {
	case stopCh <- struct{}{}: //防止阻塞，外面套一个select
	default:

	}
}
