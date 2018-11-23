package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func test()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379","1.1.1.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failded ,err: ",err)
		return
	}
	defer  cli.Close()
	for {
		rch := cli.Watch(context.Background(),"/logagent/conf/")
		for wresp := range rch {
			for _,ev := range wresp.Events{
				fmt.Printf("%s %q:%q\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
		}
	}

}
func main()  {
	test()
}
