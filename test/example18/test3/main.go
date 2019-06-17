/*
Author: lipengwei
Date: 2019/6/10
Description: 
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"path/filepath"
	"fmt"
)

type Login struct {
	User string `form: "user" json:"user" xml: "user" binding:"required"`
	Password string `form: "password json: "password xml: "password" binding: "required"`
}
func main()  {
	/*gin.DisableConsoleColor()
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"hello %s",name)
	})*/
	/*router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK,message)
	})*/
	/*router.POST("/user/:name/*action", func(c *gin.Context) {
		c.FullPath() == "/user/:name/*action"
	})*/
	/*router.POST("/somePost",posting)
	router.OPTIONS("someOptions",options)*/
	/*router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK,"helo %s %s",firstname,lastname)
		c.JSON(200,gin.H{

		})
	})
	authorized := router.Group("/")
	gin.DisableConsoleColor()
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	authorized.POST("/login",loginEndpoint)
	router.Run(":8000")*/
	/*router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\" \n",
			params.ClientIP,
			params.TimeStamp,
					params.Method,
						params.Path,
							params.Request.Proto,
								params.StatusCode,
									params.Latency,
										params.Request.UserAgent(),
											params.ErrorMessage,)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200,"pone")
	})
	router.Run(":8000")*/

	//gin.DisableConsoleColor()
	/*gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200,"pong")
	})
	router.Run(":8000")*/

	router := gin.Default()
	router.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized,gin.H{"status":"unauthorized"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"you are logged in"})
	})
	router.Run(":8000")

	/*router := gin.Default()
	router.Static("/assets","./assets")
	router.StaticFS("/more_static",http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico","./resources/favicon.ico")
	router.Run(":8000")*/
	_,file,_,_ := runtime.Caller(0)
	dir := filepath.Dir(file)
	fmt.Println(filepath.Join(dir,"/templates/"))
	router := gin.Default()
	router.LoadHTMLGlob(filepath.Join(dir,"/templates/**/*"))
	//router.LoadHTMLFiles("templates/template1.html","templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmp1",gin.H{
			"title":  "main websit",
		})
	})
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200,"/posts/index.tmp1",gin.H{
			"title" : "post",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(200,"/users/index.tpm1",gin.H{
			"title": "users",
		})
	})
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"title":"hello",
		})
	})
	router.Run(":8000")
}