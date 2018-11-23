package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	conn,err := redis.Dial("tcp","120.25.160.52:6379")
	if err != nil {
		fmt.Println("connect to redis failed ,",err.Error())
		return
	}

	defer conn.Close()
	_,err = conn.Do("lpush","book list","abc","ceg",200)
	if err != nil {
		fmt.Println("set to redis error: ",err)
		return
	}
	r,err := redis.String(conn.Do("lpop","book list"))
	if err != nil {
		fmt.Println("get abc failed , ",err)
		return
	}
	fmt.Println(r)
}
