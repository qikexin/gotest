package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func main()  {
	config := make(map[string]interface{})
	config["filename"] = "logcollect.log"
	config["level"] = logs.LevelInfo
	config["maxlinex"] = 1000000
	config["daily"] = true
	config["rotate"] = true

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed ,err: ",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr)) //相当于logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","level":"info","maxlinex":"100000","daily":"true","rotate":"true"}`)

	logs.Debug("this is a test ,my name is %s" ,"stu01")
	logs.Trace("this is a test ,my name is %s" ,"stu02")
	logs.Warn("this is a test ,my name is %s" ,"stu03")
}
