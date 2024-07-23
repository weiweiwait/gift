package gift

import (
	"gift/database"
	"gift/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	database.InitGiftInventory()
	//if err := database.ClearOrders();err != nil{
	//	panic(err)
	//}else{
	//	log.Fatalln("clear table orders")
	//}
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
