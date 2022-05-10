package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// redis 相关
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// set
	/*
		err := rdb.Set(ctx, "go_redis", "test", 0).Err()
		if err != nil {
			panic(err)
		}*/
	// get
	/*
		result, err := rdb.Get(ctx, "go_redis").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	*/
	// 获取一个不存在的key
	result, err := rdb.Get(ctx, "go_redis_v1").Result()
	if err == redis.Nil {
		fmt.Println("key 不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", result)
	}

}
