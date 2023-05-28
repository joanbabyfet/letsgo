package router

import (
	"github.com/joanbabyfet/letsgo/controller"

	"github.com/gin-gonic/gin"
)

// 路由配置, 地址与控制器对应
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")   //设置静态文件目录, 让它可外部访问到
	r.Static("/uploads", "./uploads") //设置文件上传目录, 让它可外部访问到
	r.LoadHTMLGlob("template/*")      //模板目录

	r.POST("/login", controller.Login)
	r.POST("/upload", controller.Upload)
	r.GET("/get_captcha", controller.Captcha)
	r.GET("/reload_captcha", controller.ReloadCaptcha)
	r.GET("/ping", controller.Ping)
	r.GET("/ip", controller.Ip)
	r.GET("/example", controller.ExampleIndex)
	r.GET("/example/detail", controller.ExampleDetail)
	r.POST("/example/add", controller.AddExample)
	r.POST("/example/edit", controller.EditExample)
	r.POST("/example/delete", controller.DeleteExample)
	r.POST("/example/enable", controller.EnableExample)
	r.POST("/example/disable", controller.DisableExample)
	r.GET("/example_redis", controller.ExampleRedisIndex)
	r.GET("/example_redis/detail", controller.ExampleRedisDetail)
	r.POST("/example_redis/add", controller.AddExampleRedis)
	r.POST("/example_redis/edit", controller.EditExampleRedis)
	r.POST("/example_redis/delete", controller.DeleteExampleRedis)
	r.POST("/example_redis/enable", controller.EnableExampleRedis)
	r.POST("/example_redis/disable", controller.DisableExampleRedis)
	return r
}
