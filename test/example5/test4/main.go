package main

import (
	"sync"
	"fmt"
	"time"
)

func main()  {
	var mutex sync.Mutex
	fmt.Println("lock the lock.(main)")
	mutex.Lock()
	fmt.Println("the lock is locked. (main)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("lock the lock .(g%d)\n",i)
			mutex.Lock()
			fmt.Printf("the lock is locked (g%d)\n",i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("unlock the lock.(main)")
	mutex.Unlock()
	fmt.Println("the lock is unlocked .(main)")
	time.Sleep(time.Second)
}
