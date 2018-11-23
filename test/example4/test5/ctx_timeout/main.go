package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context)  {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 1111
	}
	fmt.Printf("ret:%d\n",ret)

	sessionid, _ := ctx.Value("sessionId").(string)
	fmt.Printf("sessionid:%s\n",sessionid)

}

func main()  {
	ctx := context.WithValue(context.Background(),"trace_id",12321)
	ctx = context.WithValue(ctx,"sessionId","adfqasdcf")
	process(ctx)
}