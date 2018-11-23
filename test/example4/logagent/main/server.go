package main

import (
	"test/example4/logagent/tailf"
	"github.com/astaxie/beego/logs"
	"time"
	"test/example4/logagent/kafka"
)

func serverRun()(err error)  {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil{
			logs.Error("send to kafka failed , err :%v",err)
			time.Sleep(time.Second)
			continue
			}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg)(err error){
	//fmt.Printf("read msg:%v, topic:%v",msg.Msg,msg.Topic)
	err = kafka.SendKafka(msg.Msg,msg.Topic)
	return
}