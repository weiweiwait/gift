package gift

import (
	"gift/database"
	"gift/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	writeOrderFinish bool
)

func listenSignal() {
	c := make(chan os.Signal, syscall.SIGTERM) //注册信号,ctrl c 对应SIGINT信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-c
		if writeOrderFinish {
			log.Printf("receive signal %s, exit", sig.String())
			os.Exit(0) //进程退出
		} else {
			log.Printf("receive signal %s, but not exit", sig.String())
		}
	}
}
func Init() {
	database.InitGiftInventory()
	//if err := database.ClearOrders();err != nil{
	//	panic(err)
	//}else{
	//	log.Fatalln("clear table orders")
	//}
	handler.InitChannel()
	go func() {
		handler.TakeOrder()
		writeOrderFinish = true
	}()
	go listenSignal()
}
func main() {
	Init()
	router := gin.Default()
	router.Static("/js", "views/js")
	router.Static("/js", "views/img")
	router.StaticFile("/favicon.ico", "views/img/dpp.png")
	router.LoadHTMLFiles("views/lottery.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lottery.html", nil)
	})
	router.GET("/gifts", handler.GetAllGifts)
	router.GET("/lucky", handler.Lottery)
	router.Run("localhost:5678")
}
