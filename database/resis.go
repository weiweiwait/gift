package database

//
//const (
//	KEY_PREFIX = "auth_cookie_"
//)
//
//var (
//	blog_redis      *redis.Client
//	blog_redis_once sync.Once
//)
//
//func createRedisClient(address, passwd string, db int) *redis.Client {
//	cli := redis.NewClient(&redis.Options{
//		Addr:     address,
//		Password: passwd,
//		DB:       db,
//	})
//	if err := cli.Ping().Err(); err != nil {
//		panic(fmt.Errorf("connect to redis %d failed %v", db, err))
//
//	} else {
//		fmt.Printf("connect to redis %d\n", db)
//	}
//	return cli
//}
//func GetRedisClient() *redis.Client {
//	blog_redis_once.Do(func() {
//		if blog_redis == nil {
//			blog_redis = createRedisClient("127.0.0.1:6379", "", 0)
//		}
//	})
//	return blog_redis
//}
