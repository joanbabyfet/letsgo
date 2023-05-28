package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/global"
	"github.com/joanbabyfet/letsgo/model"
	"github.com/spf13/viper"
)

// 写入管理员操作日志
func AddAdminUserOplog(c *gin.Context, msg string, module_id int) int {
	db_name := viper.GetString("mongo.database")
	collection := global.MongoClient().Database(db_name).Collection(model.Col_admin_user_oplog)
	status := 1

	//组装数据

	//批量写入mongo
	_, err := collection.InsertMany(context.Background(), []interface{}{
		model.AdminUserOplog{
			Uid:        "1",
			Useranme:   "admin",
			Session_id: "",
			Msg:        msg,
			Module_id:  module_id,
			Op_ip:      c.ClientIP(),
			Op_country: "",
			Op_time:    int(time.Now().Unix()),
		},
	})
	if err != nil {
		fmt.Println(err)
		status = -1
	}
	return status
}
