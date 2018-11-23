package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
)

func main()  {
	c, err := redis.Dial("tcp","120.25.160.52:6379")
	if err != nil {
		fmt.Println("connect to redis failed , ",err)
		return
	}

	defer c.Close()
	key := "profile"
	imap := map[string]string{"username":"666","phoneNumber":"888"}
	value,_ := json.Marshal(imap)
	n,err := c.Do("setnx",key,value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	if n == int64(1){
		fmt.Println("success")
	}
	var imapGet map[string]string
	valuGet,err := redis.Bytes(c.Do("get",key))
	if err != nil {
		fmt.Println(err)
	}
	errshal := json.Unmarshal(valuGet,&imapGet)
	if errshal != nil {
		fmt.Println(errshal)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phoneNumber"])
}