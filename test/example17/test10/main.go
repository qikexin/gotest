/*
Author: lipengwei
Date: 2019/5/15
Description: 
*/
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
)

type User struct {
	gorm.Model
	Name string
	Role string
}
type user struct {
	ID int
	Name string
}
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name string
	Age int64
}
//type User struct {}

/*func (User) TableName() string {
	return "profiles"
}*/
func (u User)TableName() string  {
	if u.Role == "admin" {
		return "admin_users"
	}else {
		return "users"
	}
}
type USer struct {
	Name string
	Age int
}
func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200,gin.H{"html":"<b>hello,world!</b>",
		})
	})
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200,gin.H{
			"html":"<b>hello,world!</b>",
		})
	})
	r.Run(":8080")

	/*u := USer{"man",180}
	valid := validation.Validation{}
	valid.Required(u.Name,"name")
	valid.MaxSize(u.Name,15,"nameMax")
	valid.Range(u.Age,0,140,"age")
	if valid.HasErrors() {
		for _,err := range valid.Errors{
			log.Println(err.Key,err.Message)
		}
	}
	if v := valid.Max(u.Age,140,"ageMax"); !v.Ok {
		log.Println(v.Error.Key,v.Error.Message)
	}*/


	//gin.Default()默认会添加logger和recovery这两个中间件，内部调用的还是gin.new()，所以如果自定使用那些中间件，直接使用gin.new实例化一个engine对象
	/*router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login",loginEndpoint)
		v1.POST("/submit",submitEndpoint)
		v1.POST("/read",readEndpoint)
	}
	v2 := router.Group("v2")
	{
		v2.POST("/login",loginEndpoint)
		v1.POST("/submit",submitEndpoint)
		v2.POST("/read",readEndpoint)
	}
	router.Run(":8080")*/

	/*r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/benchmark",MyBenchLogger(),benchEndpoint)
	authorized := r.Group("/")
	authorized.Use(AuthRequired()) {
		authorized.POST("/login",loginEndpoint)
		authorized.POST("submit",submitEndpoint)
		authorized.POST("/read",readEndpoint)
		testing := authorized.Group("testing")
		testing.GET("analytics",analyticsEndpoint)
	}
	r.Run(":8080")*/

	/*gin.DisableConsoleColor()
	f ,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200,"pong")
	})*/

	/*router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200,"pone")
	})
	router.Run(":8080")*/




}

