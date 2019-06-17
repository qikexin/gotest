/*
Author: lipengwei
Date: 2018/12/14
Description: 
*/
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string     `json:"username"`
	NickName string		`json:"nickname"`
	Age 	 int
	Birthday string
	Sex      string
	Email    string
	Phone 	 string
}
func teststruct()(ret string,err error)  {
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age: 18,
		Birthday: "2018/8/8",
		Sex: "男",
		Email: "mahuateng@qq.com",
		Phone: "110",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json marshal failed,err",err)
		return
	}
	ret = string(data)
	return
}

func testUnMarshal()  {
	data,err := teststruct()
	if err != nil {
		fmt.Println(err)
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data),&user1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(user1)
}
func main()  {
	testUnMarshal()
}