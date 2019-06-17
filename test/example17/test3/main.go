/*
Author: lipengwei
Date: 2019/5/6
Description: 
*/
package main

import (
	"flag"
	"fmt"
	"errors"
)

var name string

func init()  {
	flag.StringVar(&name,"name","everyone","the grating object")
}

func main()  {
	flag.Parse()
	greeting,err := Hello(name)
	if err != nil {
		fmt.Printf("error: %s\n",err)
		return
	}
	fmt.Println(greeting,Introduce())
}
func Hello(name string)(string,error)  {
	if name == ""{
		return "",errors.New("empty name")
	}
	return fmt.Sprintf("hello,%s!",name),nil
}

func Introduce() string  {
	return "welcome to my golang column"
}