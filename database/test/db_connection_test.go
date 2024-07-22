package test

import (
	"gift/database"
	"testing"
)

func TestMysqlConnect(t *testing.T) {
	//db := database.GetGiftDBConnection()
	client := database.GetRedisClient()
	if client != nil {
		println(999999999999999999)
	} else {
		println(333333333)
	}
}
