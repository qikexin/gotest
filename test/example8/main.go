package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	conn, err := redis.Dial("tcp","10.9.1.210:6379")
	if err != nil {
		fmt.Println("connect to redis error,",err.Error())
	}
	defer conn.Close()
	for i:=1;i<10000;i++ {
		key := "key"+ i
		_, err = conn.Do("set",key,key)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	/*_, err = conn.Do("set","abc","100")
	if err != nil {
		fmt.Println(err)
		return
	}*/
	//r,err := redis.Int(conn.Do("get","abc"))
	//if err != nil{
	//	fmt.Println("get abc failed ",err)
	//}
	//fmt.Println(r)
}

/*func main()  {
	conn, err := redis.Dial("tcp","10.9.1.210:6379")
	if err !=nil{
		fmt.Println("connect to redis err: ",err)
		return
	}
	defer conn.Close()
	key := "profile"
	imap := map[string]string{"username":"lpw","age":"25"}
	value,_ := json.Marshal(imap)
	n,err := conn.Do("set",key,value,)
	if err != nil{
		fmt.Println(err)
		return
	}
	//fmt.Printf("success %d",n)
	if n == int64(1){  //如果发生改变，返回1，未发生改变，返回0.
		fmt.Printf("success %d",n)
	}

	var imagGet map[string]string
	valueGet,err := redis.Bytes(conn.Do("get",key))
	if err != nil {
		fmt.Println(err)
	}
	errsha1 := json.Unmarshal(valueGet,&imagGet)
	if errsha1 != nil {
		fmt.Println(errsha1)
	}
	fmt.Println(imagGet["username"])
	fmt.Println(imagGet["age"])
}*/