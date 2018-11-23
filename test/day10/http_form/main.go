package main

import (
	"net/http"
	"io"
	"fmt"
	"log"
)

const form = `<html><body><form action = "#" method = "post" name = bar>
				<input type = "text" name = "username">
				<input type = "text" name = "pwd">
				<input type = "submit" value = "Submit">
             </form></body></html>`

func SimpleServer(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello world")
}
func FormServer(w http.ResponseWriter,request *http.Request){
	w.Header().Set("Content-type","text/html")
	switch request.Method{
	case "GET":
		io.WriteString(w,form)
	case "POST":
		request.ParseForm()
		io.WriteString(w,request.FormValue("username"))
		io.WriteString(w,request.FormValue("pwd"))
	}
}
func logPanics(handle http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter,request *http.Request){
		defer func(){
			if x := recover();x != nil {
				log.Printf("[%v] caught panic: %v",request.RemoteAddr,x)
			}
		}()
		handle(writer,request)
	}
}
func main(){
	http.HandleFunc("/test1",logPanics(SimpleServer))
	http.HandleFunc("/test2",logPanics(FormServer))
	if err := http.ListenAndServe(":8088",nil);err != nil{
		fmt.Println(err)
	}
}