package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	conn,err := redis.Dial("tcp","120.25.160.52:6379")
	if err != nil {
		fmt.Println("connect to redis  failed, ",err)
		return
	}

	defer conn.Close()
	_,err = conn.Do("mset","lpw",100,"lpw1",90)
	if err != nil {
		fmt.Println(err)
		return
	}

	r,err := redis.Ints(conn.Do("mget","lpw","lpw1"))
	if err != nil {
		fmt.Println("get abc lpw failed")
		return
	}

	for _, v := range r{
		fmt.Println(v)
	}
}
