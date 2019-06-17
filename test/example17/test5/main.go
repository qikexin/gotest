/*
Author: lipengwei
Date: 2019/5/7
Description: 
*/
package main

import (
	"flag"
	"bytes"
	"sync"
	"io"
	"fmt"
	"log"
	"io/ioutil"
)

var protecting uint
func init()  {
	flag.UintVar(&protecting,"protecting",1,"it indicates whether to use a mutex to protect data writing")
}
func main()  {
	flag.Parse()
	var buffer bytes.Buffer
	const (
		max1 = 5
		max2 = 10
		max3 = 10
	)
	var mu *sync.Mutex

	sign := make(chan struct{},max1)
	for i:=1; i<=max1;i++ {
		go func(id int,writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j :=1;j<=max2;j++ {
				header := fmt.Sprintf("\n[id: %d,interation: %d]",id,j)
				data := fmt.Sprintf("%d",id * j)
				if protecting > 0 {
					mu.Lock()
				}
				_,err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]",err,id)
				}
				for k := 0;k < max3;k++{
					_,err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]",err,id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i,&buffer)
	}
	for i:=0;i<max1;i++{
		<- sign
	}
	data,err := ioutil.ReadAll(&buffer)
	if err !=nil{
		log.Fatalf("fatal error: %s",err)
	}
	log.Printf("the contents: \n%s",data)
	}