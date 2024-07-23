package mq_consumer

import (
	"context"
	"fmt"
	"gift/database"
	"github.com/bytedance/sonic"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var reader *kafka.Reader

func Init() {
	//database.InitGiftInventory()
	//
	//if err := database.ClearOrders();err != nil{
	//	panic(err)
	//}else{
	//	log.Fatalln("clear table orders")
	//}
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "topic",
		StartOffset:    kafka.LastOffset,
		GroupID:        "serialize_order",
		CommitInterval: 1 * time.Second,
	})
	log.Printf("create reader to mq")
}

//从mq里面去除订单，写入mysql

func ConsumeOrder() {
	for {
		if message, err := reader.ReadMessage(context.Background()); err != nil {
			fmt.Printf("read message from mq failed: %v,err")
			break
		} else {
			var order database.Order
			if err := sonic.Unmarshal(message.Value, &order); err == nil {
				log.Printf("message partition %d", message.Partition)
				database.CreateOrder(order.UserId, order.GiftId)
			} else {
				log.Printf("order info is invalid json format %s", string(message.Value))
			}
		}
	}
}
func listenSignal() {
	c := make(chan os.Signal, syscall.SIGTERM) //注册信号,ctrl c 对应SIGINT信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	reader.Close()
	log.Printf("receive signal %s, but not exit", sig.String())
	os.Exit(0)
}
func main() {
	Init()
	go listenSignal()
	ConsumeOrder()
}
