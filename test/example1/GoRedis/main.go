package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main(){
	client := redis.NewClient(&redis.Options{
		Addr: "120.25.160.52:6379",
		Password: "lipengwei123456",  //空字符串("")表示没有设置密码
		DB: 0,
	})
	pong,err := client.Ping().Result()
	fmt.Println(pong,err)

	err = client.Set("lpw","aaa",0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("lpw").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key",val)
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	}else if err != nil {
		panic(err)
	}else {
		fmt.Println("keys2",val2)
	}
}

