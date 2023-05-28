package global

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// 公共Db对象
var DB *gorm.DB

// 连接数据库
func InitDb() {
	driver_name := viper.GetString("db.driver_name")
	//数据来源名称, 格式如下
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
		viper.GetString("db.charset"),
	)
	database, err := gorm.Open(driver_name, dsn)
	if err != nil {
		panic("Failed to connect database, err:" + err.Error())
	}

	//指定表前缀, 操作crud会报表名错误
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	prefix := viper.GetString("db.prefix")
	// 	return prefix + defaultTableName
	// }

	DB = database
}
