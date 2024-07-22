package handler

import (
	"gift/database"
	"gift/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//获取所有奖品信息，用于初始化

func GetAllGifts(ctx *gin.Context) {
	gifts := database.GetAllGiftsV1()
	if len(gifts) == 0 {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		//抹掉敏感信息
		for _, gift := range gifts {
			gift.Count = 0
		}
		ctx.JSON(http.StatusOK, gifts)
	}
}
func Lottery(ctx *gin.Context) {
	for try := 0; try < 10; try++ {
		gifts := database.GetAllGiftInventory()
		ids := make([]int, 0, len(gifts))
		probs := make([]float64, 0, len(gifts))
		for _, gift := range gifts {
			if gift.Count > 0 {
				ids = append(ids, gift.Id)
				probs = append(probs, float64(gift.Count))
			}
		}
		if len(ids) == 0 {
			//CloseChannel
			//go CloseMQ
			ctx.String(http.StatusOK, strconv.Itoa(0)) //o表示所有奖品已经抽完
			return
		}
		index := util.Lottery(probs) //抽中第index个商品
		giftId := ids[index]
		err := database.ReduceInventory(giftId) //先从redis上减少库存
		if err != nil {
			log.Printf("奖品%d减少库存失败", giftId)
			continue
		} else {
			ctx.String(http.StatusOK, strconv.Itoa(giftId))
			return
		}

	}
	ctx.String(http.StatusOK, strconv.Itoa(database.EMPTY_GIFT)) //如果10次后还失败，返回谢谢参与
}
