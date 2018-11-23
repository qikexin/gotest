package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"test/example4/logagent/tailf"
	"time"
	"os"
)

func main()  {
	gopath := os.Getenv("GOPATH")
	filename := gopath + "/src/test/example4/logagent/conf/logagent.conf"

	//第一步加载配置
	err := loadConf("ini",filename)
	if err != nil {
		fmt.Println("load conf failed ,err:%v\n",err)
		panic("load conf failed")
		return
	}

	//初始化日志组件
	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed,err: %v\n",err)
		panic("load logger failed")
		return
	}

	logs.Debug("initialize succ")
	logs.Debug("load conf succ,config: %v",appConfig)

	err = tailf.InitTail(appConfig.collectConf,appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err: %v",err)
		return
	}
	logs.Debug("initialize all succ")

	//err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed,%v",err)
	}
	go func() {
		var count int
		for {
			count++
			logs.Debug("test for logger %d ",count)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed ,err:%v",err)
		return
	}
	logs.Debug("initialize all succ")
}
