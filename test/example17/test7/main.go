package main

import (
	"fmt"
	"net/http"
	"context"
	"time"
	"io/ioutil"
	"sync/atomic"
)
type Result struct {
	r *http.Response
	err error
	atomic.Value

}

func process()  {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second *2)
	defer  cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result,1)
	req, err:= http.NewRequest("get","http://www.baidu.com",nil)
	if err != nil {
		fmt.Println("http request failed, err",err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp,err: err}
		c <- pack
	}()
	select {
	case ctx.Done():
		tr.CancelRequest(req)
		res := <- c
		fmt.Println("timeout,err:",res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out ,_ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server response : %s",out)
	}
	return
}
func main() {
	fmt.Println("hello")
	process()
}