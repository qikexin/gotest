package main

import "fmt"

func main()  {
	dataChan := make(chan int,5)
	syncChan1 := make(chan struct{},1)
	syncChan2  := make(chan struct{},2)
	//用于演示接收操作
	go func() {
		<- syncChan1
		for {
			//ok模式用来判断当前chan是否已经被关闭（如果发送方将chan关闭，chan中未读取的元素依然可以读取）
			if elem,ok := <-dataChan;ok {
				fmt.Printf("receive: %d [receiver]\n",elem)
			}else {
				break
			}
		}
		fmt.Println("done.[received]")
		syncChan2 <- struct{}{}
	}()
	//y用于演示发送操作
	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("sent: %d [sender]\n",i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<- syncChan2
	<- syncChan2
}
