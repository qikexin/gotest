package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"test/example4/logagent/tailf"
)




type Config struct {
	logLevel 	string
	logPath  	string
	chanSize	int
	kafkaAddr   string
	collectConf []tailf.CollectConf
}

var (
	appConfig *Config
)

func loadConf(confType,filename string)(err error){
	conf, err :=config.NewConfig(confType,filename)
	if err != nil {
		fmt.Println("new config fialed,err:",err)
		return
	}

	appConfig = &Config{}

	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	appConfig.chanSize,err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed,err:%v\n",err)
		return
	}

	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr: ")
		return
	}
	return
}

func loadCollectConf(conf config.Configer)(err error){
	var cc tailf.CollectConf

	cc.Logpath = conf.String("collect::log_path")
	if len(cc.Logpath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}


	appConfig.collectConf = append(appConfig.collectConf,cc)
	return
}




























