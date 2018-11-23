package  main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)

func main()  {
	filename := "server.properties"
	tails, err := tail.TailFile(filename,tail.Config{
		ReOpen: true,
		Follow: true,
		MustExist: false,
		Poll: true,
	})
	if err != nil {
		fmt.Println("tail file err,",err)
		return
	}
	var  msg *tail.Line
	var ok bool
	for true{
		msg,ok = <- tails.Lines
		if ok {
			fmt.Printf("tail file close reopen," +
				"filename:%s\n",tails.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		fmt.Println("msg: ",msg)
	}
}