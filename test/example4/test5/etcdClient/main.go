package main

import (
	etcd_client "go.etcd.io/etcd/clientv3"
	"context"
	"fmt"
)

//func main()  {
//	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr,os.Stderr,os.Stderr))
//	cli , err := clientv3.New(clientv3.Config{
//		Endpoints: endpoints,
//		DialTimeout: dialTimeout,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cli.Close()
//	_, err = cli.Put(context.TODO(),"foo","bar")
//	if err != nil {
//		log.Fatal(err)
//	}
//}
func test()  {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			select {
			case <- ctx.Done():
				return
			case dst <- n:
				n ++
			}
		}()
		return  dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	context.TODO()
	defer  cancel()
	for n := range gen(ctx){
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
func main()  {
	cli, err := etcd_client.New(clientv3.Config{

	})
}
