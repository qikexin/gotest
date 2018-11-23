package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	conn, err := redis.Dial("tcp","120.25.160.52:6379")
	if err != nil {
		fmt.Println("connect to reids error,",err.Error())
		return
	}
	defer conn.Close()
	_,err = conn.Do("set","abc",100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r,err := redis.Int(conn.Do("get","abc"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}
	fmt.Println(r)
}