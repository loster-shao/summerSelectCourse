package model_redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	RedisPool *redis.Pool
	redisHost = "localhost:6379"
)

func newRedisPool() *redis.Pool {
	log.Println("redis start")
	return &redis.Pool{
		MaxIdle:     6000,                // 最大空闲连接数
		MaxActive:   6000,                //允许分配最大连接数
		IdleTimeout: 300 * time.Second,   // 连接时间限制
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	//创建redis连接引擎
	RedisPool = newRedisPool()
}

