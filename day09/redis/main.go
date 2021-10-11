package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() error {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis failed,err:", err)
		return
	}
	fmt.Println("连接redis成功")

	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 96, Member: "Golang"},
		redis.Z{Score: 97, Member: "Python"},
		redis.Z{Score: 98, Member: "Java"},
	}

	num, err := redisdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Println("zadd failed,err:", err)
		return
	}
	fmt.Println(num)
}
