/*
Author: lipengwei
Date: 2019/6/13
Description: 
*/
package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"net/http"
	"github.com/astaxie/beego/logs"
)
var prefix = "/data/device_report_log/"
func uploading(c *gin.Context)  {
	log := logs.NewLogger()
	log.EnableFuncCallDepth(true)
	log.Async()
	err := log.SetLogger(logs.AdapterFile,`{"filename":"/scripts/upload.log","daily":true}`)
	if err != nil {
		panic(err)
	}
	file,handler,err := c.Request.FormFile("file")
	filename := handler.Filename
	if err != nil {
		log.Error("文件 %s 上传失败!",filename)
	}
	log.Info("recieved file: %s",filename)
	out,err := os.Create(prefix+filename)
	if err != nil {
		log.Error("保存 %s 文件失败!", filename)
	}
	defer out.Close()
	_, err = io.Copy(out,file)
	if err != nil {
		log.Error("文件写入失败")
	}
	log.Info("文件[%s]上传成功!",filename)
	c.String(http.StatusOK,"upload success")

}

func main()  {
	router := gin.Default()
	router.POST("/upload",uploading)
	router.Run(":8000")
}