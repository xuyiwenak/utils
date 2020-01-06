package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	if pong, err := client.Ping().Result(); err != nil {
		fmt.Println(pong, err)
	}
	// 插入一段数据
	zAdd := client.ZAdd("ranklist", redis.Z{
		Member: "evan",
		Score:  1000,
	})
	fmt.Println(zAdd.Val())
	zscore := client.ZScore("ranklist", "evan")
	fmt.Println(zscore.Val())
	script := "local result = redis.call('ZSCORE', KEYS[1], ARGV[1])" +
		"return result"
	vals, err := client.Eval(
		script,
		[]string{"ranklist"},
		"evan",
	).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vals)
}
