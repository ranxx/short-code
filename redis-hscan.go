package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

//InitRedis 初始化redis
func InitRedis(server, pwd string) error {
	option := &redis.Options{
		Addr:     server,
		Password: pwd,
	}
	redisClient = redis.NewClient(option)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	// 启动docker redis
	err := InitRedis("127.0.0.1:6379", "")
	if err != nil {
		panic(err)
	}
	write()
}

func write() {
	muser := map[string]interface{}{}
	for i := 0; i < 2048; i++ {
		muser[fmt.Sprintf("key-%v", i+10)] = fmt.Sprintf("value-%v", i+10)
	}
	stat := redisClient.HMSet("user", muser)
	fmt.Println(stat.Result())
}

func read() {
	cursor, count := uint64(0), int64(5)
	for {
		scan := redisClient.HScan("user", cursor, "", count)
		keys, _newCusor, err := scan.Result()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if _newCusor == 0 {
			break
		}
		cursor = _newCusor
		fmt.Printf("%#v\n", keys)
		// 通过HMGET获取user keys 的value
	}
}

func main() {
	read()
}
