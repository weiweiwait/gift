package handler

import (
	"gift/database"
	"github.com/gin-gonic/gin"
)

func Lottery(ctx *gin.Context) {
	for try := 0; try < 10; try++ {
		gifts := database.GetAllGiftsV1()
		println(gifts)
	}
}
