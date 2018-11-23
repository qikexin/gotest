package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	c,err := redis.Dial("tcp","120.25.160.52:6379")
	if err != nil {
		fmt.Println("connect redis failed, ",err.Error())
		return
	}
	defer c.Close()
	_,err = c.Do("hset","books","abc",90,)
	if err != nil {
		fmt.Println(err)
		return
	}
	r,err := redis.Int(c.Do("hget","books","abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
