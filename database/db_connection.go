package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type StringContextKey string

var (
	blog_mysql      *gorm.DB
	blog_mysql_once sync.Once
	dblog           ormlog.Interface
	blog_redis      *redis.Client
	blog_redis_once sync.Once
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      ormlog.Info,
			Colorful:      true,
		},
	)
}

func createMysqlDB(dbname, host, user, pass string, port int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, pass, host, port, dbname)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true})
	if err != nil {
		//util.LogRus.Panicf("connect to mysql use dsn %s failed: %s", dsn, err)
		log.Fatalf("connect to mysql use dsn %s failed: %s", dsn, err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	//util.LogRus.Infof("connect to mysql db %s", dbname)
	log.Printf("connect to mysql db %s", dbname)
	return db
}
func GetGiftDBConnection() *gorm.DB {
	blog_mysql_once.Do(func() {
		blog_mysql = createMysqlDB("gift", "localhost", "root", "Fjw20030504", 3306)
	})

	return blog_mysql
}

func createRedisClient(address, passwd string, db int) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: passwd,
		DB:       db,
	})
	if err := cli.Ping().Err(); err != nil {
		panic(fmt.Errorf("connect to redis %d failed %v", db, err))

	} else {
		fmt.Printf("connect to redis %d\n", db)
	}
	return cli
}
func GetRedisClient() *redis.Client {
	blog_redis_once.Do(func() {
		if blog_redis == nil {
			blog_redis = createRedisClient("127.0.0.1:6379", "", 0)
		}
	})
	return blog_redis
}
