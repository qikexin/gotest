package tailf

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)
type CollectConf struct {
	Logpath string
	Topic string
}

//TailObj代表一个被监控的日志文件对象
type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}
type TextMsg struct {
	Msg  string
	Topic string
}


type TailObjMgr struct {
	tailObjs []*TailObj	   //TailObjMgr 代表多个被监控的日志文件对象，所以tailObjs是数组类型。
	msgChan chan *TextMsg
}

var tailObjMgr *TailObjMgr

func GetOneLine()(msg *TextMsg){
	msg =  <- tailObjMgr.msgChan
	return
}
func InitTail(conf []CollectConf,chanSize int)(err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect,conf: %v",conf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg,chanSize),
	}
	for _,v := range conf {
		obj := &TailObj{
			conf: v,
		}
		tails, errTail := tail.TailFile(v.Logpath,tail.Config{
			ReOpen: true,
			Follow: true,
			MustExist: false,
			Poll: true,
		})
		if  errTail != nil {
			fmt.Println("tail file err: ",err)
		}
		obj.tail = tails

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs,obj)
		go readFromTail(obj)
	}

	return
}

func readFromTail(tailObj *TailObj)  {
	for true {
		line, ok := <- tailObj.tail.Lines
		if !ok {

			logs.Warn("tail file close reopen,filename: %s\n",tailObj.tail.Filename)
			//fmt.Printf("tail file close reopen,filename: %s\n",tails.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg : line.Text,
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}
}