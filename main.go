package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/joanbabyfet/letsgo/config"
	"github.com/joanbabyfet/letsgo/global"
	"github.com/joanbabyfet/letsgo/model"
	"github.com/joanbabyfet/letsgo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	//设置日志保存到日志文件, 按照日期保存
	date := time.Now().Format("2006-01-02")
	log_file := path.Join("./logs", fmt.Sprintf("gin-%s.log", date))
	file, err := os.OpenFile(log_file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file) //设置存储位置
}

func main() {
	//加载路由文件
	r := router.InitRouter()
	//加载配置文件
	config.InitConfig()
	//创建数据库连接
	global.InitDb()

	//生成table
	global.DB.Debug().AutoMigrate(&model.Example{})

	//defer关键字，当api访问结束时，关闭数据库连接
	defer global.DB.Close()

	//监听8080端口
	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}
