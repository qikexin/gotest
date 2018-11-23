package main

import "fmt"

//func createworker(id int) chan int {
//	c := make(chan int)
//	go func() {
//		for  {
//			fmt.Printf("worker %d received %d\n",id,<-c)
//		}
//	}()
//	return c
//}

//func chanDemo(){
//	//c := make(chan int)
//	var channels [10] chan int
//	for i := 0; i < 10; i++ {
//		channels[i] = createworker(i)
//
//	}
//	for i := 0 ; i < 10; i++ {
//		channels[i] <-  'a' +i
//	}
//	time.Sleep(time.Millisecond)
//}
func chanClose(){
	c := make(chan  int)
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
func readchan(id int ,c chan int){
	for n := range c{
		fmt.Printf("%d received %d",id,n)
	}

}
func main()  {
	//chanDemo()
	chanClose()
	readchan()
}
