/*
Author: lipengwei
Date: 2019/5/30
Description: 
*/
package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
)

type User struct {
	Name string `json: "name"`
	Age  int	`json: "age"`
}

func index(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println("Form: ",r.Form)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println(r.Form["a"])
	fmt.Println(r.Form["b"])
	for k, v := range r.Form {
		fmt.Println(k, "=>", v, strings.Join(v,"-"))
	}
	fmt.Fprintf(w,"it works !")
}

func test(w http.ResponseWriter, r *http.Request)  {
	body,_ := ioutil.ReadAll(r.Body)
	body_str := string(body)
	fmt.Println(body_str)
	var user User
	if err := json.Unmarshal(body,&user);err == nil {
		fmt.Println(user)
		user.Age += 10
		fmt.Println(user)
		ret,_:= json.Marshal(user)
		fmt.Fprint(w,string(ret))
	}else {
		fmt.Println(err)
	}
}
func main()  {
	http.HandleFunc("/",index)
	http.HandleFunc("/test",test)
	if err := http.ListenAndServe(":8080",nil);err != nil {
		log.Fatal("listinAndServer:", err)
	}
}