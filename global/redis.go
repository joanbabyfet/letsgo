package global

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type connect struct {
	client *redis.Client
}

var once = sync.Once{}
var _connect *connect

// 连接redis
func connectRedis() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	addr := viper.GetString("redis.host") + ":" + viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	db, _ := strconv.Atoi(viper.GetString("redis.db"))
	pool_size, _ := strconv.Atoi(viper.GetString("redis.poolsize"))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: pool_size,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错, 错误信息：%v", err)
	}
	//fmt.Println(pong)

	_connect = &connect{
		client: client,
	}
}

// 使用单例模式, 提供一个外部访问方法 类似 GetInstance()
func RedisClient() *redis.Client {
	if _connect == nil {
		once.Do(func() {
			connectRedis()
		})
	}
	return _connect.client
}
