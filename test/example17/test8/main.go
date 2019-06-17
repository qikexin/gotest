/*
Author: lipengwei
Date: 2019/5/15
Description: 
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"context"
)

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("the number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 0; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j <= stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("end")
}
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP,currNum,newNum){
			fmt.Printf("the number: %d [%d-%d]\n",numP,id,i)
			break
		}
	}
}
type myKey int
func main()  {
	//coordinateWithWaitGroup()
	keys := []myKey{
		myKey(20),
		myKey(30),
		myKey(60),
		myKey(61),
	}
	values := []string{
		"value in node2",
		"value in node3",
		"value in node6",
		"value in node6Branch",
	}
	rootNode := context.Background()
	node1,cancelFunc1 := context.WithCancel(rootNode)
	defer cancelFunc1()

	node2 := context.WithValue(node1,keys[0],values[0])
	node3 := context.WithValue(node2,keys[1],values[1])
	fmt.Printf("the value of the key %v found in the node3: %v\n",keys[0],node3.Value(keys[0]))
	fmt.Printf("the value of the key %v found in the node3: %v\n",keys[1],node3.Value(keys[1]))
	fmt.Printf("the value of the key %v found in the node3: %v\n",keys[2],node3.Value(keys[2]))
	fmt.Println()

	node4,_ :=context.WithCancel(node3)
	node5,_ := context.WithTimeout(node4,time.Hour)
	fmt.Printf("the value of the key %v found in the node5: %v\n",keys[0],node5.Value(keys[0]))
	fmt.Printf("the value of the key %v found in the node5: %v\n",keys[1],node5.Value(keys[1]))
	fmt.Println()
}