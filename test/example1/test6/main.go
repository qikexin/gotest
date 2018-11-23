package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var pool *redis.Pool
func init(){
	pool = &redis.Pool{
		MaxIdle: 16,
		MaxActive: 0,
		IdleTimeout: 300,
		Dial: func()(redis.Conn,error){
			return redis.Dial("tcp","120.25.160.52:6379")
		},
	}
}
func main()  {
	c := pool.Get()
	defer c.Close()
	if _,err := c.Do("auth","lipengwei123456");err != nil {
		c.Close()
		fmt.Println(err)
		return
	}
	_, err := c.Do("set","abc",100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r,err := redis.Int(c.Do("get","abc"))
	if err != nil {
		fmt.Println("get abc failed ,", err)
		return
	}
	fmt.Println(r)
	pool.Close()
}
