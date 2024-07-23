package handler

import (
	"context"
	"gift/database"
	"github.com/bytedance/sonic"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

var (
	writer    *kafka.Writer
	writeWg   sync.WaitGroup
	closeOnce sync.Once
)

func InitMQ() {
	writer = &kafka.Writer{
		Addr:                   kafka.TCP("127.0.0.1:9092"),
		Topic:                  "topic",
		AllowAutoTopicCreation: true,
	}
	log.Printf("create writer to mq")
}

//把订单放入mq

func ProduceOrder(UserId, GiftId int) {
	order := database.Order{
		UserId: UserId,
		GiftId: GiftId,
	}
	writeWg.Add(1)
	go func() { //写mq太慢,异步执行
		defer writeWg.Done()
		json, _ := sonic.Marshal(order)
		if err := writer.WriteMessages(context.Background(), kafka.Message{Value: json}); err != nil {
			log.Fatalf("write kafka failed %s", err)
		}
	}()
}

//关闭mq,closeMq可以被反复调用

func CloseMQ() {
	closeOnce.Do(func() {
		writeWg.Wait() //因为写mq诗异步执行的，要等所有写操作完了，才能关闭writer
		writer.Close()
		log.Printf("stop writer mq")
	})
}
