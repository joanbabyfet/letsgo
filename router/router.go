package router

import (
	"github.com/joanbabyfet/letsgo/controller"
	"github.com/joanbabyfet/letsgo/middleware"

	"github.com/gin-gonic/gin"
)

// 路由配置, 地址与控制器对应
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")   //设置静态文件目录, 让它可外部访问到
	r.Static("/uploads", "./uploads") //设置文件上传目录, 让它可外部访问到
	r.LoadHTMLGlob("template/*")      //模板目录

	//跨域中间件
	r.Use(middleware.Cors())

	r.GET("/ping", controller.Ping)
	r.GET("/ip", controller.Ip)
	r.POST("/login", controller.Login)
	r.GET("/get_captcha", controller.Captcha)
	r.GET("/reload_captcha", controller.ReloadCaptcha)

	// 使用jwt中间件
	authorized := r.Group("/")
	authorized.Use(middleware.JwtAuth())
	{
		r.POST("/upload", controller.Upload)
		authorized.GET("/example", controller.ExampleIndex)
		authorized.GET("/example/detail", controller.ExampleDetail)
		authorized.POST("/example/add", controller.AddExample)
		authorized.POST("/example/edit", controller.EditExample)
		authorized.POST("/example/delete", controller.DeleteExample)
		authorized.POST("/example/enable", controller.EnableExample)
		authorized.POST("/example/disable", controller.DisableExample)
		authorized.GET("/example_redis", controller.ExampleRedisIndex)
		authorized.GET("/example_redis/detail", controller.ExampleRedisDetail)
		authorized.POST("/example_redis/add", controller.AddExampleRedis)
		authorized.POST("/example_redis/edit", controller.EditExampleRedis)
		authorized.POST("/example_redis/delete", controller.DeleteExampleRedis)
		authorized.POST("/example_redis/enable", controller.EnableExampleRedis)
		authorized.POST("/example_redis/disable", controller.DisableExampleRedis)
	}
	return r
}
