package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func main()  {
	config := make(map[string]interface{})
	config["filename"] = "logcollect.og"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed , err",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	logs.Debug("this is a test debug,")
	logs.Trace("this is a trace")
	logs.Warn("this is a warn")
}
